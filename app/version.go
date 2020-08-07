package app

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
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

//var URL = ""
var GIT_SHELL_NAME = "git"
var GIT_REMOTE = "origin"
var GIT_MAIN_BRANCH = "master"

func SetGitShellName(name string) {
	GIT_SHELL_NAME = name
}

func SetGitRemote(name string) {
	GIT_REMOTE = name
}

func SetGitMainBranch(name string) {
	GIT_MAIN_BRANCH = name
}

/**
 * 获取当前版本数据
 */
func GetCommitLogInfo(branch string) []Commit {
	logs, startTime := getDiffLog(branch)
	fetchRemoteBranch(branch, GIT_REMOTE)
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
	shell := exec.Command(GIT_SHELL_NAME, "log", "..."+branch, `--format=%h`+split+`%ci`+split+`%ce`+split+`%s`, "--no-merges")
	output, err := shell.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", common, err.Error())
		return nil, ""
	}
	log.Println("output:", string(output))
	if string(output) == "" {
		return nil, ""
	}
	logList := strings.Split(string(output), "\n")
	var GitLogSlice = make([]Commit, 0)
	//log.Println("log list:", logList)
	var stratTime string
	for _, log_string := range logList {
		var GitLog Commit
		logStringList := strings.Split(string(log_string), split)
		if len(logStringList) != 4 {
			log.Println("continue")
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

func GetProjectNameAndRemoteUrl() (string, string) {
	shell := exec.Command(GIT_SHELL_NAME, "remote", "get-url", GIT_REMOTE)
	output, err := shell.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", "", err.Error())
		//return nil, ""
	}
	gitUrl := string(output)
	log.Println("output:", gitUrl)
	if gitUrl == "" {
		return "", ""
	}
	gitUrl = formatGitUrl(gitUrl)
	s := strings.Split(gitUrl, "/")
	var le int = len(s) - 1
	var projectName string = s[le]
	projectName = strings.Replace(projectName, ".git", "", 1)
	VersionInfo.ProjectGitName = projectName
	VersionInfo.ProjectGitUrl = gitUrl
	return gitUrl, projectName
}

func formatGitUrl(gitUrl string) string {
	gitUrl = strings.Replace(gitUrl, "\n", "", 1)
	//gitUrl = strings.
	t := strings.Split(gitUrl, "@")
	if len(t) > 1 {
		log.Println(t[0][0:3])
		switch {
		case len(t[0]) >= 5 && t[0][0:5] == "https":
			gitUrl = "https://" + t[1]
			break
		case len(t[0]) >= 4 && t[0][0:4] == "http":
			gitUrl = "http://" + t[1]
			break
		case len(t[1]) >= 10 && t[1][0:10] == "github.com":
			gitUrl = "https://" + strings.Replace(t[1], ":", "/", 1)
			break
		case t[0] == "git":
			// TODO : 提供项目配置读取方法
			log.Println("ssh 协议不能自动获取到正确的http的git地址，后续会提供配置支持，敬请期待")
			//gitUrl = "http://" + strings.Replace(t[1], ":", "/", 1)
			break
		default:
			log.Println("unknown protocol :" + gitUrl)
			gitUrl = "http://" + t[1]
			break
		}
	}
	return gitUrl
}

func GetChangeFileInfo(branch string) string {
	shell := exec.Command(GIT_SHELL_NAME, "diff", ".."+branch, "--stat")
	output, err := shell.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", "", err.Error())
		return ""
	}
	diffLogStat := string(output)
	VersionInfo.CodeChangeFileInfo = diffLogStat
	return diffLogStat
}

func fetchRemoteBranch(branch, remote string) {
	common := `git fetch ` + remote + " " + branch
	shell := exec.Command(GIT_SHELL_NAME, "fetch", remote, branch)
	output, err := shell.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", common, err.Error())
		//return nil, ""
		return
	}
	log.Println(output)
}

func getBranchCommitHashs(branch string, startTime string, ) map[string]bool {
	// 拉取远程分支到本地

	common := `git log ` + branch
	//git log dev --after="2020-08-03 00:00:00" --format=%H
	shell := exec.Command(GIT_SHELL_NAME, "log", branch, `--after="`+startTime+`"`, `--format=%H`)
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
	url := VersionInfo.ProjectGitUrl
	url = strings.Replace(url, ".git", "/commit/", 1)
	for _, commit := range gitLog {
		commitLog = commitLog + fmt.Sprintf("%d. %s [%s](%s%s)\n", i, commit.Desc, commit.Hash, url, commit.Hash)
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

func CreateVersionFile(version string) {
	branch := GIT_MAIN_BRANCH
	GetProjectNameAndRemoteUrl()

	VersionInfo.Version = version
	GetFormatCommitLog(branch)
	GetChangeFileInfo(branch)
	content := GetTemplateContent()
	VersionInfo.DateTime = time.Now().Format("2006-01-02 15:04:05")
	content = ReplaceContent(content, VersionInfo)
	CreateFile(content, version)
}
