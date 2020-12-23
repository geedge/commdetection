/*
 * @Author: your name
 * @Date: 2020-12-22 16:36:18
 * @LastEditTime: 2020-12-22 16:36:44
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/cmd/version_test.go
 */

package cmd

import "testing"

func BenchmarkVersion(b *testing.B) {
	version()
}
