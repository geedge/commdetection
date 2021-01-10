/*
 * @Author: EnderByEndera
 * @Date: 2021-01-06 11:12:49
 * @LastEditTime: 2021-01-10 15:33:10
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/model/db.go
 */

package model

import (
	"commdetection/logger"
	"commdetection/yaml"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbset   yaml.MongoSet
	timeout time.Duration
)

func init() {
	dbset = yaml.GetMongoSetting()
	var err error
	timeout, err = time.ParseDuration(dbset.Timeout)
	if err != nil {
		logger.Warnln(err)
	}
}

func closeClient(ctx context.Context, client *mongo.Client) {
	err := client.Disconnect(ctx)
	if err != nil {
		logger.Warnln(err)
	}
}

func connect(ctx context.Context) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d", dbset.Host, dbset.Port)))
}
