/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"F10-CLI/app"
	"fmt"
	"github.com/spf13/cobra"
)

var version string
var branch string
var remote string
var overwrite bool = false

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "生成版本描述文件",
	Long: `
generate命令将对比当前分支与主分支（默认为master，可通过--branch=指定）差异，提取commit log 生成更新日志

使用示例：
	F10-CLI generate 版本(任务)编号
	
    指定对比分支 默认与master对比
	F10-CLI generate 版本(任务)编号 --branch master --remote origin 

    强制重新生成 默认为false
	F10-CLI generate 版本(任务)编号 --overwrite

`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version = args[0]
		if len(branch) == 0 {
			branch = "master"
		}

		if len(remote) == 0 {
			branch = "origin"
		}
		//
		fmt.Println("即将生成" + version + "版本描述文件")
		// 判断版本文件是否存在 存在则退出
		if overwrite == false {
			if app.IsExist("./version/" + version + ".md") {
				fmt.Println("./version/" + version + ".md 文件已存在")
				return
			}
		}
		app.CreateVersionFile(version, branch, remote)

		fmt.Println("版本描述文件已生成")
		// 获取版本数据

		// 生成文件

		// fmt.Println("")

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&branch, "branch", "b", "master", "对比分支")
	generateCmd.Flags().StringVarP(&remote, "remote", "r", "origin", "远程地址别名")
	generateCmd.Flags().BoolVarP(&overwrite, "overwrite", "o", false, "强制覆盖重复版本")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
