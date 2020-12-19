/*
 * @Author: EnderByEndera
 * @Date: 2020-12-16 13:59:51
 * @LastEditTime: 2020-12-19 10:43:36
 * @LastEditors: Please set LastEditors
 * @Description: Includes CommScore definition
 * @FilePath: /commdetection/rules/commscore.go
 */

package rules

import (
	"commdetection/comm"
)

// CommScore includes command name and its score
type CommScore struct {
	Command comm.Command `json:"command"`
	Score   float64      `json:"score"`
}

// InitCommScores initialize commscores from []Command
func InitCommScores(commands []comm.Command) (commScores []CommScore) {
	for _, command := range commands {
		commScores = append(commScores, CommScore{
			Command: command,
			Score:   100,
		})
	}
	return
}
