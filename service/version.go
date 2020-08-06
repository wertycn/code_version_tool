package service

import "time"

type VersionType struct {
	Version    string
	Project    string
	CreateTime string
	UpdateTime string
	UserName   string
}

var versionMap = make(map[string]VersionType)

func IsExistVersion(version string) bool {
	_, ok := versionMap[version]
	return ok
}

func RegVersion(version, project, username string) VersionType {
	var versionInfo VersionType
	versionInfo.Version = version
	versionInfo.Project = project
	versionInfo.UserName = username
	versionInfo.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	if IsExistVersion(version) {
		versionInfo.CreateTime = versionMap[version].CreateTime
	} else {
		versionInfo.CreateTime = versionInfo.UpdateTime
	}
	versionMap[version] = versionInfo
	go saveVersionMap()
	return versionInfo
}



// TODO:已发布版本信息写入文本，服务重启后读取文本信息
// TODO:异步保存版本信息到本地文件
func saveVersionMap() {

}

// TODO:加载本地版本
func LoadLocalVersionMap() {

}

