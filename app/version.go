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
func GetVersionInfo(branch string) []GIT_LOG {
	logs, startTime := getDiffLog(branch)
	hashMap := getBranchCommitHashs(branch, startTime)
	var GitLogSlice []GIT_LOG
	for _, GitLog := range logs {
		//fmt.Println(i)
		if _, ok := hashMap[GitLog.Hash]; ok {
			fmt.Println("ok", ok)
			fmt.Println("ok GitLog", GitLog)
			continue
		}
		fmt.Println("GitLog:", GitLog)
		GitLogSlice = append(GitLogSlice, GitLog)
	}
	return GitLogSlice
}

/**
 * 获取差异日志对象及最早时间
 */
func getDiffLog(branch string) ([]GIT_LOG, string) {
	split := " -|---|- "
	common := `git log ...` + branch + `  --format="%H` + split + `%ci` + split + `%ce` + split + `%s"`
	//common := "git version"
	shell := exec.Command("git", "log", "..."+branch, `--format=%H`+split+`%ci`+split+`%ce`+split+`%s`)
	output, err := shell.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", common, err.Error())
		return nil, ""
	}
	fmt.Println("output:", string(output))
	if string(output) == "" {
		return nil, ""
	}
	log_list := strings.Split(string(output), "\n")
	var GitLogSlice = make([]GIT_LOG, 10)
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
		GitLogSlice = append(GitLogSlice, GitLog)
		stratTime = GitLog.Date
	}
	return GitLogSlice, stratTime

}

func getBranchCommitHashs(branch string, startTime string) map[string]int {
	common := `git log ` + branch
	//git log dev --after="2020-08-03 00:00:00" --format=%H
	shell := exec.Command("git", "log", branch, `--after="`+startTime+`"`, `--format=%H`)
	output, err := shell.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", common, err.Error())
		//return nil, ""
		return nil
	}
	hash_list := strings.Split(string(output), "\n")
	var hash_list_res = make(map[string]int)
	for _, hash := range hash_list {
		if hash != "" {
			hash_list_res[hash] = 1
		}
	}
	return hash_list_res
}
