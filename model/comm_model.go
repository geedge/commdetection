/*
 * @Author: your name
 * @Date: 2021-01-06 09:56:18
 * @LastEditTime: 2021-01-19 12:02:23
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/model/comm_model.go
 */

package model

import (
	"commdetection/logger"
	"fmt"
	"reflect"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//IntSlice includes a list of n array in order to use sort.Sort() method to sort n
type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] > s[j] }

// Command contains command and its flags or symbols
type Command struct {
	CommName  string    `json:"commname" bson:"commname"`
	Args      []string  `json:"args" bson:"args,omitempty"`
	Flags     []string  `json:"flags" bson:"flags,omitempty"`
	TimeStamp time.Time `json:"timestamp" bson:"timestamp"`
	User      string    `json:"user" bson:"user"`
	Mac       string    `json:"mac" bson:"mac"`
}

// Commands is the multi type of Command
type Commands []Command

func (c Commands) Len() int {
	return len(c)
}

func (c Commands) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Commands) Less(i, j int) bool {
	return c[i].TimeStamp.Before(c[j].TimeStamp)
}

// Has returns whether c has the command
func (c Commands) Has(command Command) bool {
	sort.Sort(c)
	if sort.Search(c.Len(), func(i int) bool {
		return reflect.DeepEqual(c[i], command)
	}) == c.Len() {
		return false
	}
	return true
}

// removeOneCommand removes only one command from the list given the index
func removeOneCommand(commands []Command, n int) []Command {
	return append(commands[:n], commands[n+1:]...)
}

// RemoveCommands removes a list of commands using index
func RemoveCommands(commands []Command, n IntSlice) []Command {
	if len(n) == 0 {
		return commands
	}
	sort.Sort(n)
	for i := range n {
		commands = removeOneCommand(commands, n[i])
	}
	return commands
}

// GetCommandsFrom gets all the commands in the mongodb collections
func (c *Commands) GetCommandsFrom(dbName string, cName string) error {
	return mongoOpsWithoutIndex(getCommandsFromFn, opParams{
		dbName:   dbName,
		cName:    cName,
		commands: c,
	})
}

func getCommandsFromFn(sc mongo.SessionContext) error {
	client := sc.Client()
	params, ok := sc.Value(key("params")).(opParams)
	if !ok {
		return fmt.Errorf("Error transfering the params")
	}
	collection := client.Database(params.dbName).Collection(params.cName)
	cur, err := collection.Find(sc, bson.D{})
	if err != nil {
		return err
	}
	defer cur.Close(sc)
	for cur.Next(sc) {
		var next Command
		err := cur.Decode(&next)
		if err != nil {
			logger.Warnln(err)
		}
		*params.commands = append(*params.commands, next)
	}
	return nil
}

// InsertAllTo insert the given commands to the specified database and collection
func (c *Commands) InsertAllTo(dbName string, cName string) error {
	return mongoOpsWithoutIndex(insertAllCommandsToFn, opParams{
		dbName:   dbName,
		cName:    cName,
		commands: c,
	})
}

func insertAllCommandsToFn(sc mongo.SessionContext) error {
	client := sc.Client()
	params, ok := sc.Value(key("params")).(opParams)
	if !ok {
		return fmt.Errorf("Error tranfering the params")
	}
	if params.commands.Len() == 0 {
		return nil
	}

	var deleteIndexes []int
	for i := 0; i < params.commands.Len()-1; i++ {
		if reflect.DeepEqual((*params.commands)[i], (*params.commands)[i+1]) {
			deleteIndexes = append(deleteIndexes, i)
		}
	}
	*params.commands = RemoveCommands(*params.commands, deleteIndexes)
	collections := client.Database(params.dbName).Collection(params.cName)
	var documents []interface{}
	for _, command := range *params.commands {
		// if the command is not found in the mongodb
		if res := collections.FindOne(sc, command); res.Err() == mongo.ErrNoDocuments {
			// add the command to the ready-inserted documents
			documents = append(documents, command)
		}
	}
	// If every command is inserted into the db, documents may be nil
	if documents == nil {
		return nil
	}
	_, err := collections.InsertMany(sc, documents)
	if err != nil {
		return err
	}
	return nil
}

// InsertAnyTo inserts one command to the dbName.cName
func (c *Commands) InsertAnyTo(dbName, cName string, index uint) error {
	return mongoOpsWithIndex(insertAnyCommandToFn, opParams{
		dbName:   dbName,
		cName:    cName,
		index:    index,
		commands: c,
	})
}

func insertAnyCommandToFn(sc mongo.SessionContext) error {
	client := sc.Client()
	params, ok := sc.Value(key("params")).(opParams)
	if !ok {
		return fmt.Errorf("Error transfering the params")
	}
	collections := client.Database(params.dbName).Collection(params.cName)
	if res := collections.FindOne(sc, (*params.commands)[int(params.index)]); res.Err() != mongo.ErrNoDocuments {
		return nil
	}
	_, err := collections.InsertOne(sc, (*params.commands)[int(params.index)])
	if err != nil {
		return err
	}
	return nil
}

// UpdateAnyTo updates the command in the mongodb
func (c *Commands) UpdateAnyTo(dbName, cName string, index uint, updateFilter interface{}) error {
	return mongoOpsWithIndex(updateAnyCommandFn, opParams{
		dbName:       dbName,
		cName:        cName,
		index:        index,
		commands:     c,
		updateFilter: updateFilter,
	})
}

func updateAnyCommandFn(sc mongo.SessionContext) error {
	client := sc.Client()
	params, ok := sc.Value(key("params")).(opParams)
	if !ok {
		return fmt.Errorf("Error transfering the params")
	}
	collection := client.Database(params.dbName).Collection(params.cName)
	command := (*params.commands)[int(params.index)]
	// if the command is found in the mongodb, update will be useless, so return nil
	if res := collection.FindOne(sc, command); res.Err() != mongo.ErrNoDocuments {
		return nil
	}
	_, err := collection.UpdateOne(sc, params.updateFilter, bson.D{{
		"$set",
		command,
	}})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneFrom deletes one command from the dbName.cName
func (c *Commands) DeleteOneFrom(dbName, cName string, index uint) error {
	return mongoOpsWithIndex(deleteOneCommandFromFn, opParams{
		dbName:   dbName,
		cName:    cName,
		index:    index,
		commands: c,
	})
}

func deleteOneCommandFromFn(sc mongo.SessionContext) error {
	client := sc.Client()
	params, ok := sc.Value(key("params")).(opParams)
	if !ok {
		return fmt.Errorf("Error transfering the params")
	}
	collections := client.Database(params.dbName).Collection(params.cName)
	_, err := collections.DeleteOne(sc, (*params.commands)[int(params.index)])
	if err != nil {
		return err
	}
	return nil
}

// DeleteAllFrom deletes many commands from dbName.cName
func (c *Commands) DeleteAllFrom(dbName, cName string) error {
	return mongoOpsWithoutIndex(deleteAllCommandsFromFn, opParams{
		dbName:   dbName,
		cName:    cName,
		commands: c,
	})
}

func deleteAllCommandsFromFn(sc mongo.SessionContext) error {
	client := sc.Client()
	params, ok := sc.Value(key("params")).(opParams)
	if !ok {
		return fmt.Errorf("Error tranfering the params")
	}
	collections := client.Database(params.dbName).Collection(params.cName)
	var deleteResults []*mongo.DeleteResult
	for _, command := range *params.commands {
		res, err := collections.DeleteOne(sc, command)
		if err != nil {
			return err
		}
		deleteResults = append(deleteResults, res)
	}
	return nil
}
