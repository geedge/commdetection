/*
 * @Author: EnderByEndera
 * @Date: 2021-01-06 15:26:48
 * @LastEditTime: 2021-02-01 07:03:10
 * @LastEditors: Please set LastEditors
 * @Description: test mongo.go
 * @FilePath: /commdetection/model/mongo_test.go
 */

package model

import (
	"reflect"
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
		now := time.Now()
		for i := 0; i < 10; i++ {
			commands = append(commands, Command{
				CommName:  "wget",
				Args:      []string{},
				Flags:     []string{},
				TimeStamp: now,
				User:      "root",
				Mac:       "25:21:4A:1A:4C:11",
			})
		}
		err := commands.InsertAllTo("test", "commands")
		if err != nil {
			t.Error(err)
		}
		commands.InsertAllTo("test", "commands")
		if err != nil {
			t.Error(err)
		}
		commands.GetCommandsFrom("test", "commands")
		for i := 0; i < commands.Len()-1; i++ {
			if reflect.DeepEqual(commands[i], commands[i+1]) {
				t.Error("commands has same data")
			}
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
	_ = commands.InsertAnyTo("test", "commands", 0)

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
		err := commands.InsertAllTo("test", "commands")
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
