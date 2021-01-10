/*
 * @Author: EnderByEndera
 * @Date: 2020-12-16 13:59:51
 * @LastEditTime: 2021-01-06 09:58:52
 * @LastEditors: Please set LastEditors
 * @Description: Includes CommScore definition
 * @FilePath: /commdetection/rules/commscore.go
 */

package rules

import (
	"commdetection/model"
)

// InitCommScores initialize commscores from []Command
func InitCommScores(commands []model.Command) (commScores []model.CommScore) {
	for _, command := range commands {
		commScores = append(commScores, model.CommScore{
			Command: command,
			Score:   100,
		})
	}
	return
}
