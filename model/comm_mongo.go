/*
 * @Author: EnderByEndera
 * @Date: 2021-01-09 16:46:53
 * @LastEditTime: 2021-01-10 16:31:17
 * @LastEditors: Please set LastEditors
 * @Description: comm_mongo.go tries to get connection with mongodb lib
 * @FilePath: /commdetection/model/command_mongo.go
 */

package model

import (
	"commdetection/logger"
	"context"
	"fmt"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetCommandsFrom gets all the commands in the mongodb collections
func (c *Commands) GetCommandsFrom(dbName string, cName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, err := connect(ctx)
	if err != nil {
		return err
	}
	defer closeClient(ctx, client)
	return client.UseSession(ctx, func(sc mongo.SessionContext) error {
		client := sc.Client()
		collection := client.Database(dbName).Collection(cName)
		cur, err := collection.Find(ctx, bson.D{})
		defer cur.Close(ctx)
		if err != nil {
			return err
		}
		for cur.Next(ctx) {
			var next Command
			err := cur.Decode(&next)
			if err != nil {
				logger.Warnln(err)
			}
			*c = append(*c, next)
		}
		return nil
	})
}

// InsertAllTo insert the given commands to the specified database and collection
func (c *Commands) InsertAllTo(dbName string, cName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, err := connect(ctx)
	if err != nil {
		return err
	}
	if !sort.IsSorted(c) {
		sort.Sort(c)
	}
	defer closeClient(ctx, client)
	return client.UseSession(ctx, func(sc mongo.SessionContext) error {
		client := sc.Client()
		collections := client.Database(dbName).Collection(cName)
		var documents []interface{}
		for _, command := range *c {
			documents = append(documents, command)
		}
		res, err := collections.InsertMany(sc, documents, options.InsertMany().SetOrdered(true))
		if err != nil {
			sc.AbortTransaction(ctx)
			return err
		}
		sc.CommitTransaction(ctx)
		logger.Debugln("Insert succeeded, Ids are ", res)
		return nil
	})
}

// InsertAnyTo inserts one command to the dbName.cName
func (c *Commands) InsertAnyTo(dbName, cName string, index uint) error {
	if index >= uint(c.Len()) {
		return fmt.Errorf("Index %d out of range", index)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, err := connect(ctx)
	if err != nil {
		return err
	}
	defer closeClient(ctx, client)
	return client.UseSession(ctx, func(sc mongo.SessionContext) error {
		client := sc.Client()
		collections := client.Database(dbName).Collection(cName)
		res, err := collections.InsertOne(sc, (*c)[index])
		if err != nil {
			sc.AbortTransaction(sc)
			return err
		}
		sc.CommitTransaction(sc)
		logger.Debugln("Insert succeeded, Id is ", res)
		return nil
	})
}

// DeleteOneFrom deletes one command from the dbName.cName
func (c *Commands) DeleteOneFrom(dbName, cName string, index uint) error {
	if index >= uint(c.Len()) {
		return fmt.Errorf("index %d out of range", index)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, err := connect(ctx)
	if err != nil {
		return err
	}
	defer closeClient(ctx, client)
	return client.UseSession(ctx, func(sc mongo.SessionContext) error {
		client := sc.Client()
		collections := client.Database(dbName).Collection(cName)
		res, err := collections.DeleteOne(sc, (*c)[index])
		if err != nil {
			sc.AbortTransaction(ctx)
			return err
		}
		sc.CommitTransaction(ctx)
		logger.Debugln("Delete succeeded, Id is ", res)
		return nil
	})
}

// DeleteAllFrom deletes many commands from dbName.cName
func (c *Commands) DeleteAllFrom(dbName, cName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, err := connect(ctx)
	if err != nil {
		return err
	}
	defer closeClient(ctx, client)
	return client.UseSession(ctx, func(sc mongo.SessionContext) error {
		client := sc.Client()
		collections := client.Database(dbName).Collection(cName)
		var deleteResults []*mongo.DeleteResult
		for _, command := range *c {
			res, err := collections.DeleteOne(ctx, command)
			if err != nil {
				sc.AbortTransaction(ctx)
				return err
			}
			deleteResults = append(deleteResults, res)
		}
		sc.CommitTransaction(ctx)
		logger.Debugln("Delete succeeded, ids are ", deleteResults)
		return nil
	})
}
