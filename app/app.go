package app

import (
	"F10-CLI/util"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//import _ "os"
var PROJCET_VERSION_PATH = "./version"

type USER_RUN_INFO_TYPE struct {
	LatestPullTime time.Time
	ProjectPath    string
	Remote         string
	MainBranch     string
}

/**
 * 创建文件
 */
func CreateFile(content string, fileName string) {

	file2, error := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(file2)
	defer file2.Close()
	file2.WriteString(content)
}

func IsVersionPathExistAndCreate() {
	if IsExist(PROJCET_VERSION_PATH) == false {
		os.Mkdir(PROJCET_VERSION_PATH, os.ModePerm)
	}
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func GetTemplateContent() string {
	home := os.Getenv("F10_CLI_HOME")
	fmt.Println(home)
	f, err := os.Open(home + "./template/version.md")
	defer f.Close()
	if err != nil {
		fmt.Println("文件读取错误")
		return ""
	}
	contentByte, e := ioutil.ReadAll(f)
	if e != nil {
		return ""
	}
	return string(contentByte)
}

func ReplaceContent(content string, versionInfo VersionInfoType) string {
	content = strings.ReplaceAll(content, "{{ DATETIME }}", versionInfo.DateTime)
	content = strings.ReplaceAll(content, "{{ DATE }}", versionInfo.DateTime)
	content = strings.ReplaceAll(content, "{{ PROJECT_GIT_URL }}", versionInfo.ProjectGitUrl)
	content = strings.ReplaceAll(content, "{{ PROJECT_GIT_NAME }}", versionInfo.ProjectGitName)
	content = strings.ReplaceAll(content, "{{ TASK_NO }}", versionInfo.Version)
	content = strings.ReplaceAll(content, "{{ DEVELOPMENT }}", versionInfo.Development)
	content = strings.ReplaceAll(content, "{{ CODE_CHANGE_FILE_COUNT }}", versionInfo.CodeChangeFileCount)
	content = strings.ReplaceAll(content, "{{ CODE_CHANGE_FILE_INFO }}", versionInfo.CodeChangeFileInfo)
	content = strings.ReplaceAll(content, "{{ CODE_CHANGE_COMMIT_INFO }}", versionInfo.CodeChangeCommitInfo)
	return content
}

func IsPullMainBranch() bool {
	//	判断项目上次提交主分支的时间
	now := time.Now().Unix()
	latestPull := GetLatestPullTime()
	i := now - latestPull
	fmt.Println(i)
	if i > 36000 {
		return true
	}
	return false
}

// 获取项目根目录
func GetMainDir() (string, error) {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	sep := string(os.PathSeparator)
	dirSlice := strings.Split(pwd, sep)
	fmt.Println(sep)
	gitDir := ""
	for i := 0; i < len(dirSlice); i++ {
		if IsExist(gitDir + ".git") {
			projectPath, _ := filepath.Abs(pwd + sep + gitDir)
			return projectPath, nil
		}
		gitDir = "../" + gitDir
	}

	return "", errors.New("没有找到git项目根目录")
}

func SavePullTime() {
	pullTimeConfigPath := GetPullTimeConfigPath()
	if IsExist(pullTimeConfigPath) == false {
		os.Mkdir(pullTimeConfigPath, os.ModePerm)
	}
	file, error := os.OpenFile(pullTimeConfigPath+"latest_pull_time", os.O_CREATE, 0766)
	defer file.Close()
	if error != nil {
		log.Println("保存项目pull时间失败：" + error.Error())
		return
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	file.WriteString(timestamp)
	log.Println("更新项目pull主分支时间成功")
}

func GetLatestPullTime() int64 {
	pullTimeConfigPath := GetPullTimeConfigPath()
	if IsExist(pullTimeConfigPath + "latest_pull_time") {
		return 0
	}
	content, err := ioutil.ReadFile(pullTimeConfigPath + "latest_pull_time")
	if err != nil {
		return 0
	}
	timestamp := string(content)
	return util.StrTime2Int(timestamp)
}

func GetPullTimeConfigPath() string {
	projectPath, err := GetMainDir()
	if err != nil {
		panic(err.Error())
	}
	sep := string(os.PathSeparator)
	pullTimeConfigPath := projectPath + sep + ".F10-CLI" + sep
	return pullTimeConfigPath
}
