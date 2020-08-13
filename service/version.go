package service

import (
	"F10-CLI/app"
	"bytes"
	"encoding/json"
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type VersionType struct {
	Version    string   `json:version`
	Project    string   `json:project`
	CreateTime string   `json:create_time`
	UpdateTime string   `json:update_time`
	Service    []string `json:service`
}

var versionMap = make(map[string]VersionType)
var appCommitIdMap = make(map[string]VersionType)
var VERSION_ARCHIVE_NAME = "./version_archive.json"

func SetVersionArchiveFileName(name string) {
	VERSION_ARCHIVE_NAME = name
}

func IsExistVersion(version string) bool {
	_, ok := versionMap[version]
	return ok
}

func IsExistApp(version, app string) bool {
	if IsExistVersion(version) {
		versionType := versionMap[version]
		for _, alreadyService := range versionType.Service {
			if app == alreadyService {
				return true
			}
		}
	}
	return false
}

func IsExistCommitId(app, commitId string) bool {
	_, ok := appCommitIdMap[app+"_"+commitId]
	return ok
}

func RegVersion(version, project, app, commitId string) VersionType {
	var versionInfo VersionType
	versionInfo.Version = version
	versionInfo.Project = project
	versionInfo.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	if IsExistVersion(version) {
		versionInfo.CreateTime = versionMap[version].CreateTime
		// 判断app是否存在
		exist := false
		for _, _app := range versionMap[version].Service {
			if _app == app {
				exist = true
			}
		}
		if false == exist {
			versionInfo.Service = append(versionMap[version].Service, app)
		} else {
			versionInfo.Service = versionMap[version].Service
		}
	} else {
		versionInfo.Service = []string{app}
		versionInfo.CreateTime = versionInfo.UpdateTime
	}
	versionMap[version] = versionInfo
	if commitId != "" {
		appCommitIdMap[app+"_"+commitId] = versionInfo
	}
	return versionInfo
}

func UnRegVersion(version, project, app, commitId string) VersionType {
	var versionInfo VersionType
	if IsExistVersion(version) {
		versionInfo = versionMap[version]
		delete(versionMap, version)
		delete(versionMap, version)
	}
	if commitId != "" && IsExistCommitId(app, commitId) {
		delete(appCommitIdMap, app+"_"+commitId)
	}
	go SaveVersionMap()
	return versionInfo
}

// TODO:已发布版本信息写入文本，服务重启后读取文本信息
// TODO:异步保存版本信息到本地文件
func SaveVersionMap() {
	data, err := json.Marshal(versionMap)
	if err != nil {
		log.Println("持久化版本数据失败：" + err.Error())
		return
	}
	file, error := os.OpenFile(VERSION_ARCHIVE_NAME, os.O_CREATE, 0766)
	defer file.Close()
	if error != nil {
		log.Println("持久化版本数据失败：" + error.Error())
		return
	}
	file.WriteString(string(data))
	log.Println("持久化版本数据成功")
	log.Println(string(data))

}

// TODO:加载本地版本
func LoadLocalVersionMap() {
	if false == app.IsExist(VERSION_ARCHIVE_NAME) {
		log.Println("加载持久化版本数据失败：" + VERSION_ARCHIVE_NAME + "文件不存在。")
		return
	}
	file, err := os.Open(VERSION_ARCHIVE_NAME)
	defer file.Close()
	//file, err := JsonLoad(VERSION_ARCHIVE_NAME)
	var _versionMap = make(map[string]VersionType)
	if err != nil {
		log.Println("加载持久化版本数据失败：" + err.Error())
		// 重新创建
		//versionMap = make(map[string]VersionType)
		return
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&_versionMap)
	if err != nil {
		log.Println("decoder failed")
	}
	versionMap = _versionMap
	//versionMap = file.(map[string]VersionType)
	log.Println("加载持久化版本数据成功")
}

func JsonLoad(filename string) (*simplejson.Json, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil
	}
	return simplejson.NewJson(data)
}

func UnMarshalJson(req *http.Request, v interface{}) error {
	result, err := ioutil.ReadAll(req.Body)
	fmt.Println(req)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(bytes.NewBuffer(result).String()), v)
	return nil
}
