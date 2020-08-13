/*
Copyright © 2020 DEBUG.ICU <debugicu@163.com>

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

var commitType string
var commitSubject string
var noAdd bool = false
var noPull bool = false

var commitTypeDoc string = `type: 
    feat: 新特性，新功能
    fix: 修复问题
    refactor: 代码重构
    docs: 文档修改
    style: 代码格式修改, 注意不是 css 修改
    test: 测试用例修改
    pref: 性能提升的修改
    build: 对项目构建或者依赖的改动
    ci: CI 的修改
    revert: revert 前一个 commit
    chore: 其他修改, 比如构建流程, 依赖管理等`

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "提交代码",
	Long: `规范commit提交，提交前会自动执行git add all ， 提交后如果超过1天未pull主分支，会自动执行git pull origin master操作，可以使用--no-add 和 --no-pull参数禁用自动行为。

F10 commit规范是基于被业界广泛认可的 [Angular commit message规范]的极简版，之所以精简，是为了减少提交过程中的阻碍，同时，鼓励大家在开发过程中保持小步快跑的方式，在程序可工作的状态下，频繁进行提交，而不是采用整个项目开发完成后才提交了一次代码。
<type>: <subject>
类型: 对修改内容的简要描述（中文不得少于8个字,英文不得少于24个字母）
` + commitTypeDoc + `
subject: 
    commit 主题，对修改内容的简要描述，不得少于8个字`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		commitType = args[0]
		commitSubject = args[1]
		if false == app.CheckCommitType(commitType) {
			fmt.Println(`请输入正确的commit type`)
			fmt.Println(commitTypeDoc)
			fmt.Println("提交失败！")
			return
		}
		fmt.Println("提交成功~")
		if false == app.CheckCommitSubjectLength(commitSubject) {
			fmt.Println("检查commit内容不通过，请至少输入24个字符，一个中文占3个字符")
			return
		}
		if noAdd == false {
			app.GitAddAll()
		}
		app.SubmitCommit(commitSubject)

		if noPull == false && app.IsPullMainBranch() {
			fmt.Println("满足pull主分支条件，即将执行git pull " + remote + " " + branch)
			app.GitPull(remote, branch)
			app.SavePullTime()
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringVarP(&branch, "branch", "b", "master", "主分支名 默认master")
	commitCmd.Flags().StringVarP(&remote, "remote", "r", "origin", "远程地址名 默认origin")
	commitCmd.Flags().BoolVarP(&noAdd, "no-add", "", false, "本次提交不执行git add --all")
	commitCmd.Flags().BoolVarP(&noPull, "no-pull", "", false, "本次提交后不执行git pull")
}
