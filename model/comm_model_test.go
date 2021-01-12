/*
 * @Author: EnderByEndera
 * @Date: 2021-01-06 15:26:48
 * @LastEditTime: 2021-01-12 16:49:55
 * @LastEditors: Please set LastEditors
 * @Description: test mongo.go
 * @FilePath: /commdetection/model/mongo_test.go
 */

package model

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func TestGetCommandsFromMongo(t *testing.T) {
	commands := Commands{}
	err := commands.GetCommandsFrom("test", "commands")
	if err != nil {
		t.Error(err)
	}
	if commands == nil || commands.Len() == 0 {
		t.Error("Commands are nil commands")
	}
}

func TestInsertCommandsToMongo(t *testing.T) {
	t.Run("Test InsertManyTo func", func(t *testing.T) {
		commands := Commands{}
		for i := 0; i < 10; i++ {
			commands = append(commands, Command{
				CommName:  "wget",
				Args:      []string{},
				Flags:     []string{},
				TimeStamp: time.Now(),
				User:      "root",
			})
		}
		err := commands.InsertAllTo("test", "commands")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Test InsertAnyTo func", func(t *testing.T) {
		commands := Commands{
			{
				CommName:  "wget",
				Args:      []string{},
				Flags:     []string{},
				TimeStamp: time.Now(),
				User:      "root",
			},
		}
		err := commands.InsertAnyTo("test", "commands", 1)
		if err == nil {
			t.Errorf("No IndexOutOfRange error appeared")
		}
		err = commands.InsertAnyTo("test", "commands", 0)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestUpdateCommandsToMongo(t *testing.T) {
	commands := Commands{
		{
			CommName:  "vim",
			Args:      []string{},
			Flags:     []string{},
			TimeStamp: time.Now(),
			User:      "root",
		},
	}
	err := commands.UpdateAnyTo("test", "commands", 0, bson.D{{"commname", "wget"}})
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteCommandsFromMongo(t *testing.T) {
	t.Run("Test DeleteOne func", func(t *testing.T) {
		commands := Commands{
			{
				CommName:  "wget",
				Args:      []string{},
				Flags:     []string{},
				TimeStamp: time.Now(),
				User:      "root",
			},
		}
		err := commands.InsertAnyTo("test", "commands", 0)
		if err != nil {
			t.Error(err)
		}
		err = commands.DeleteOneFrom("test", "commands", 0)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Test DeleteAll func", func(t *testing.T) {
		commands := Commands{}
		for i := 0; i < 10; i++ {
			commands = append(commands, Command{
				CommName:  "wget",
				Args:      []string{},
				Flags:     []string{},
				TimeStamp: time.Now(),
				User:      "root",
			})
		}
		err := commands.InsertAllTo("test", "commands")
		if err != nil {
			t.Error(err)
		}
		err = commands.DeleteAllFrom("test", "commands")
		if err != nil {
			t.Error(err)
		}
	})
}
