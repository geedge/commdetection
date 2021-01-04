/*
 * @Author: EnderByEndera
 * @Date: 2020-12-16 13:59:51
 * @LastEditTime: 2021-01-04 16:41:41
 * @LastEditors: Please set LastEditors
 * @Description: Includes CommScore definition
 * @FilePath: /commdetection/rules/commscore.go
 */

package rules

import (
	"commdetection/comm"
	"commdetection/model"
)

// InitCommScores initialize commscores from []Command
func InitCommScores(commands []comm.Command) (commScores []model.CommScore) {
	for _, command := range commands {
		commScores = append(commScores, model.CommScore{
			Command: command,
			Score:   100,
		})
	}
	return
}
