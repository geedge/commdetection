/*
 * @Author: your name
 * @Date: 2021-01-06 09:56:18
 * @LastEditTime: 2021-01-10 09:45:28
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/model/comm_model.go
 */

package model

import (
	"reflect"
	"sort"
	"time"
)

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
	return c[j].TimeStamp.After(c[i].TimeStamp)
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
