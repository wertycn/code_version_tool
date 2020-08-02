package app

import (
	"fmt"
	"os/exec"
	"strings"
)

type GIT_LOG struct {
	Hash   string
	Date   string
	Author string
	Commit string
}

/**
 * 获取当前版本数据
 */
func GetVersionInfo() {
	logs, startTime := getDiffLog()
	fmt.Println(logs)
	fmt.Println(startTime)
}

/**
 * 获取差异日志对象及最早时间
 */
func getDiffLog() (map[string]GIT_LOG, string) {
	split := " -|---|- "
	common := `git log ...master  --format="%H` + split + `%ci` + split + `%ce` + split + `%s"`
	//common := "git version"
	shell := exec.Command("git", "log", `--format=%H`+split+`%ci`+split+`%ce`+split+`%s`)
	output, err := shell.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", common, err.Error())
		return nil, ""
	}
	log_list := strings.Split(string(output), "\n")
	GitLogMap := make(map[string]GIT_LOG)
	fmt.Println("log list:", log_list)
	var stratTime string
	for _, log_string := range log_list {
		fmt.Println("log_string", log_string)
		//fmt.Println("log_string", n)

		var GitLog GIT_LOG
		log_string_list := strings.Split(string(log_string), split)
		if len(log_string_list) != 4 {
			fmt.Println(log_string)
			continue
		}

		GitLog.Hash = log_string_list[0]
		GitLog.Date = log_string_list[1][0:19]
		GitLog.Author = log_string_list[2]
		GitLog.Commit = log_string_list[3]
		GitLogMap[GitLog.Hash] = GitLog
		stratTime = GitLog.Date
	}
	return GitLogMap, stratTime

}

func getBranchCommitHashs(branch string,startTime string) {
	common := `git log ` + branch
	shell := exec.Command("git", "log", `--format=%H`+split+`%ci`+split+`%ce`+split+`%s`)
	output, err := shell.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", common, err.Error())
		return nil, ""
	}
	log_list := strings.Split(string(output), "\n")
}
