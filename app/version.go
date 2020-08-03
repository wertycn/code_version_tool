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

type VersionInfo struct {
	Version              string
	ProjectGitName       string
	ProjectGitUrl        string
	Development          string
	DateTime             string
	CodeChangeFileCount  string
	CodeChangeFileInfo   string
	CodeChangeCommitInfo string
}

/**
 * 获取当前版本数据
 */
func GetCommitLogInfo(branch string) []GIT_LOG {
	logs, startTime := getDiffLog(branch)
	hashMap := getBranchCommitHashs(branch, startTime)
	var GitLogSlice []GIT_LOG
	for _, GitLog := range logs {
		if _, ok := hashMap[GitLog.Hash]; ok {
			continue
		}
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
	logList := strings.Split(string(output), "\n")
	var GitLogSlice = make([]GIT_LOG, 0)
	//fmt.Println("log list:", logList)
	var stratTime string
	for _, log_string := range logList {
		var GitLog GIT_LOG
		logStringList := strings.Split(string(log_string), split)
		if len(logStringList) != 4 {
			fmt.Println("continue")
			continue
		}
		GitLog.Hash = logStringList[0]
		GitLog.Date = logStringList[1][0:19]
		GitLog.Author = logStringList[2]
		GitLog.Commit = logStringList[3]
		GitLogSlice = append(GitLogSlice, GitLog)
		stratTime = GitLog.Date
	}
	return GitLogSlice, stratTime

}

func getBranchCommitHashs(branch string, startTime string) map[string]bool {
	common := `git log ` + branch
	//git log dev --after="2020-08-03 00:00:00" --format=%H
	shell := exec.Command("git", "log", branch, `--after="`+startTime+`"`, `--format=%H`)
	output, err := shell.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", common, err.Error())
		//return nil, ""
		return nil
	}
	hashList := strings.Split(string(output), "\n")
	var hashListRes = make(map[string]bool)
	for _, hash := range hashList {
		if hash != "" {
			hashListRes[hash] = true
		}
	}
	return hashListRes
}
