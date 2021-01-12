/*
 * @Author: EnderByEndera
 * @Date: 2021-01-04 16:30:53
 * @LastEditTime: 2021-01-12 11:53:39
 * @LastEditors: Please set LastEditors
 * @Description: This is the model file used for rules pack
 * @FilePath: /commdetection/model/rule_model.go
 */

package model

import (
	"commdetection/logger"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Rule defines a rule's func and its name
type Rule struct {
	Name     string `json:"name"`
	RuleFunc string `json:"rulefunc"`
}

// Rules is the slice of Rule
type Rules []Rule

// Evaluation is the template of the Evaluation function
type Evaluation func(CommScore) CommScore

// CommScore includes command name and its score
type CommScore struct {
	Command Command `json:"command" bson:"command"`
	Score   float64 `json:"score" bson:"score"`
}

// CommScores is the multiple type of CommScore
type CommScores []CommScore

func (css CommScores) Len() int {
	return len(css)
}

func (css CommScores) Swap(i, j int) {
	css[i], css[j] = css[j], css[i]
}

func (css CommScores) Less(i, j int) bool {
	return css[i].Command.TimeStamp.Before(css[j].Command.TimeStamp)
}

// SPath includes sensitive path dir and its sensitive coefficient
type SPath struct {
	Path        string
	Coefficient float64
}

// SPaths is the SPath's multiple type
type SPaths []SPath

// SComm includes sensitive command's name and its sensitive coefficient
type SComm struct {
	Comm        string  `json:"command"`
	Coefficient float64 `json:"coefficient"`
}

// SComms is the SComm's multiple types
type SComms []SComm

// Ussites is the Unsensitive websites, a.k.a website whitelists
// Coefficient defines the intent of lowering the score if the website arg is not in the ussites.Websites
type Ussites struct {
	Websites    []string `json:"websites"`
	Coefficient float64  `json:"coefficient"`
}

// GetCommScoresFrom gets command scores from dbName.cName collection
func (css *CommScores) GetCommScoresFrom(dbName, cName string) error {
	return mongoOpsWithoutIndex(getCommScoreFromFn, opParams{
		dbName: dbName,
		cName:  cName,
		css:    css,
	})
}

func getCommScoreFromFn(sc mongo.SessionContext) error {
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
	for cur.Next(sc) {
		var next CommScore
		err := cur.Decode(&next)
		if err != nil {
			return err
		}
		*params.css = append(*params.css, next)
	}
	return nil
}

// InsertAllTo insert command scores to the dbName.cName collection
func (css *CommScores) InsertAllTo(dbName, cName string) error {
	return mongoOpsWithoutIndex(insertAllCommScoreToFn, opParams{
		dbName: dbName,
		cName:  cName,
		css:    css,
	})
}

func insertAllCommScoreToFn(sc mongo.SessionContext) error {
	client := sc.Client()
	params, ok := sc.Value(key("params")).(opParams)
	if !ok {
		return fmt.Errorf("Error transfering the params")
	}
	collection := client.Database(params.dbName).Collection(params.cName)
	documents := []interface{}{}
	for _, cs := range *params.css {
		documents = append(documents, cs)
	}
	res, err := collection.InsertMany(sc, documents)
	if err != nil {
		return err
	}
	logger.Debugln("insert succeeded, ids are ", res)
	return nil
}

// InsertAnyTo inserts one command score to the mongodb database
func (css *CommScores) InsertAnyTo(dbName, cName string, index uint) error {
	return mongoOpsWithIndex(insertAnyCommScoreToFn, opParams{
		dbName: dbName,
		cName:  cName,
		index:  index,
		css:    css,
	})
}

func insertAnyCommScoreToFn(sc mongo.SessionContext) error {
	client := sc.Client()
	params, ok := sc.Value(key("params")).(opParams)
	if !ok {
		return fmt.Errorf("Error transfering the params")
	}
	collection := client.Database(params.dbName).Collection(params.cName)
	res, err := collection.InsertOne(sc, (*params.css)[int(params.index)])
	if err != nil {
		return err
	}
	logger.Debugln("insert succeeded, id is ", res)
	return nil
}

// UpdateAnyTo updates the command score in the mongodb
func (css *CommScores) UpdateAnyTo(dbName, cName string, index uint, updateFilter interface{}) error {
	return mongoOpsWithIndex(updateAnyCommScoreToFn, opParams{
		dbName:       dbName,
		cName:        cName,
		css:          css,
		index:        index,
		updateFilter: updateFilter,
	})
}

func updateAnyCommScoreToFn(sc mongo.SessionContext) error {
	client := sc.Client()
	params, ok := sc.Value(key("params")).(opParams)
	if !ok {
		return fmt.Errorf("Error transfering the params")
	}
	collection := client.Database(params.dbName).Collection(params.cName)
	res, err := collection.UpdateOne(sc, params.updateFilter, (*params.css)[int(params.index)])
	if err != nil {
		return err
	}
	logger.Debugln("update succeeded, id is ", res)
	return nil
}

// DeleteOneFrom deletes one command score in the mongodb
func (css *CommScores) DeleteOneFrom(dbName, cName string, index uint) error {
	return mongoOpsWithIndex(deleteOneCommScoreFromFn, opParams{
		dbName: dbName,
		cName:  cName,
		index:  index,
		css:    css,
	})
}

func deleteOneCommScoreFromFn(sc mongo.SessionContext) error {
	client := sc.Client()
	params, ok := sc.Value(key("params")).(opParams)
	if !ok {
		return fmt.Errorf("Error transfering the params")
	}
	collection := client.Database(params.dbName).Collection(params.cName)
	res, err := collection.DeleteOne(sc, (*params.css)[int(params.index)])
	if err != nil {
		return err
	}
	logger.Debugln("delete succeeded, id is ", res)
	return nil
}

// DeleteAllFrom deletes all command scores in the mongodb which are the same as css
func (css *CommScores) DeleteAllFrom(dbName, cName string) error {
	return mongoOpsWithoutIndex(deleteAllCommScoresFromFn, opParams{
		dbName: dbName,
		cName:  cName,
		css:    css,
	})
}

func deleteAllCommScoresFromFn(sc mongo.SessionContext) error {
	client := sc.Client()
	params, ok := sc.Value(key("params")).(opParams)
	if !ok {
		return fmt.Errorf("Error transfering the params")
	}
	collection := client.Database(params.dbName).Collection(params.cName)
	deleteResults := []*mongo.DeleteResult{}
	for _, cs := range *params.css {
		res, err := collection.DeleteOne(sc, cs)
		if err != nil {
			return err
		}
		deleteResults = append(deleteResults, res)
	}
	logger.Debugln("delete succeeded, id are ", deleteResults)
	return nil
}
