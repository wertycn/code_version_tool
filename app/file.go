package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//import _ "os"

/**
 * 创建文件
 */
func CreateFile(content string, title string) {
	// 判断目录是否存在 不存在则创建
	if IsExist("./version") == false {
		os.Mkdir("./version", os.ModePerm)
	}

	fileName := "./version/" + title + ".md"
	file2, error := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(file2)
	defer file2.Close()
	file2.WriteString(content)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func GetTemplateContent() string {
	f, err := os.Open("../template/version.md")
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
	content = strings.ReplaceAll(content, "{{ TASK_NO }}", versionInfo.Version)
	content = strings.ReplaceAll(content, "{{ DEVELOPMENT }}", versionInfo.Development)
	content = strings.ReplaceAll(content, "{{ CODE_CHANGE_FILE_COUNT }}", versionInfo.CodeChangeFileCount)
	content = strings.ReplaceAll(content, "{{ CODE_CHANGE_FILE_INFO }}", versionInfo.CodeChangeFileInfo)
	content = strings.ReplaceAll(content, "{{ CODE_CHANGE_COMMIT_INFO }}", versionInfo.CodeChangeCommitInfo)
	return content
}
