package app

import (
	"fmt"
	"os/exec"
	"strings"
)

type Commit struct {
	Hash   string
	Date   string
	Author string
	Desc   string
}

type CommitChange struct {
	Feat     string
	Fix      string
	Docs     string
	Style    string
	Refactor string
	Perf     string
	Test     string
	Revert   string
	Other    string
}

type VersionInfoType struct {
	Version              string
	ProjectGitName       string
	ProjectGitUrl        string
	Development          string
	DateTime             string
	CodeChangeFileCount  string
	CodeChangeFileInfo   string
	CodeChangeCommitInfo string
}

var VersionInfo VersionInfoType
var URL = ""

/**
 * 获取当前版本数据
 */
func GetCommitLogInfo(branch string) []Commit {
	logs, startTime := getDiffLog(branch)
	hashMap := getBranchCommitHashs(branch, startTime)
	var GitLogSlice []Commit
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
func getDiffLog(branch string) ([]Commit, string) {
	split := " -|---|- "
	common := `git log ...` + branch + `  --format="%h` + split + `%ci` + split + `%ce` + split + `%s"`
	//common := "git version"
	shell := exec.Command("git", "log", "..."+branch, `--format=%h`+split+`%ci`+split+`%ce`+split+`%s`)
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
	var GitLogSlice = make([]Commit, 0)
	//fmt.Println("log list:", logList)
	var stratTime string
	for _, log_string := range logList {
		var GitLog Commit
		logStringList := strings.Split(string(log_string), split)
		if len(logStringList) != 4 {
			fmt.Println("continue")
			continue
		}
		GitLog.Hash = logStringList[0]
		GitLog.Date = logStringList[1][0:19]
		GitLog.Author = logStringList[2]
		GitLog.Desc = logStringList[3]
		GitLogSlice = append(GitLogSlice, GitLog)
		stratTime = GitLog.Date
	}
	return GitLogSlice, stratTime

}

func GetProjectNameAndRemoteUrl(name string) (string, string) {
	shell := exec.Command("git", "remote", "get-url", name)
	output, err := shell.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", "", err.Error())
		//return nil, ""
	}
	gitUrl := string(output)
	fmt.Println("output:", gitUrl)
	if gitUrl == "" {
		return "", ""
	}
	gitUrl = strings.Replace(gitUrl, "\n", "", 1)

	s := strings.Split(gitUrl, "/")
	var le int = len(s) - 1
	var projectName string = s[le]
	projectName = strings.Replace(projectName, ".git", "", 1)
	VersionInfo.ProjectGitName = projectName
	VersionInfo.ProjectGitUrl = gitUrl
	return gitUrl, projectName
}

func GetChangeFileInfo(branch string) string {
	shell := exec.Command("git", "diff", ".."+branch, "--stat")
	output, err := shell.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", "", err.Error())
		return ""
	}
	diffLogStat := string(output)
	VersionInfo.CodeChangeFileInfo = diffLogStat
	return diffLogStat
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

func FormatCommitLog(gitLog []Commit) (string, string) {
	var commitLog string
	var authorMap = make(map[string]string)
	var i = 1
	for _, commit := range gitLog {
		commitLog = commitLog + fmt.Sprintf("%d. %s [%s](%s)\n", i, commit.Desc, commit.Hash, URL)
		//commitLog = commitLog + string(i) + "." + commit.Desc + " [" + commit.Hash + "](" + URL + ")\n"
		i = i + 1
		if _, ok := authorMap[commit.Author]; ok {
			continue
		}

		authorMap[commit.Author] = commit.Author
	}
	var authors string
	for author, _ := range authorMap {
		authors = authors + author + "; "
	}
	return commitLog, authors
}

func GetFormatCommitLog(branch string) (string, string) {
	info := GetCommitLogInfo(branch)
	commit, authors := FormatCommitLog(info)
	VersionInfo.CodeChangeCommitInfo = commit
	VersionInfo.Development = authors
	return commit, authors
}

func CreateVersionFile(version string, branch string,remote string) {
	VersionInfo.Version = version
	GetFormatCommitLog(branch)
	GetChangeFileInfo(branch)
	GetProjectNameAndRemoteUrl(remote)
	content := GetTemplateContent()
	content = ReplaceContent(content, VersionInfo)
	CreateFile(content, version)
}
