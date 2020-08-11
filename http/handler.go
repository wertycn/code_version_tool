package http

import (
	"F10-CLI/app"
	"F10-CLI/service"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

const (
	uploadPath string = "./upload/"
)

type UploadData struct {
	ImageUrl     string      `json:"image_url"`
	CropImageUrl string      `json:"crop_image_url"`
	CropStatus   bool        `json:"crop_status"`
	FaceInfo     interface{} `json:"face_info"`
}

func handlerTest(writer http.ResponseWriter, request *http.Request) {

}

func handlerUploadRequest(writer http.ResponseWriter, request *http.Request) {
	// 获取上传的图片
	resp := &Resp{Code: 0}
	defer MarshalJson(writer, &resp)
	request.ParseForm()
	form := request.Form
	version := GetParam(form, "version", "").(string)
	project := GetParam(form, "project", "").(string)
	appName := GetParam(form, "app", "").(string)
	commitId := GetParam(form, "commit_id", "").(string)
	repeat := GetParam(form, "repeat", "").(string)
	if version == "" || project == "" || appName == "" {
		resp.Code = -1
		resp.Msg = "必要参数（version，project，service）不能为空"
		return
	}
	if service.IsExistCommitId(appName, commitId) && commitId != "" {
		resp.Code = -1
		resp.Msg = "该版本已存在发布记录，不允许重复发布，如确实需要重新发布，请勾选重新发布"
		return
	}
	if service.IsExistApp(version, appName) && repeat != "1" {
		resp.Code = -1
		resp.Msg = "该应用版本已发布，不允许重复发布，如发布失败重新发布，请勾选重新发布"
		return
	}
	//curl "https://httpbin.org/post" -H "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:79.0) Gecko/20100101 Firefox/79.0" -H "Accept: application/json, text/javascript, */*; q=0.01" -H "Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2" --compressed -H "Content-Type: multipart/form-data; boundary=---------------------------32936956637119493282264586759" -H "Origin: https://www.layui.com" -H "Connection: keep-alive" -H "Referer: https://www.layui.com/demo/upload.html" -H "Pragma: no-cache" -H "Cache-Control: no-cache" --data-raw ""
	//curl  -F "file=@\"D:/dev/F10-CLI/app/version/LOCAL_TEST.md\"" "http://localhost:9980/upload?version=F10-336&project=f10-cli&username=hanjinxiang"
	file, head, err := request.FormFile("file")
	if err != nil {
		log.Println("文件上传失败", err)
		resp.Msg = "文件上传失败：" + err.Error()
		resp.Code = -1
		return
	}
	defer file.Close()
	if path.Ext(head.Filename) != ".md" {
		resp.Code = -1
		resp.Msg = "版本描述文件必须为markdown文件"
		return
	}
	saveFileName := version + path.Ext(head.Filename)
	saveDir := uploadPath + project + (time.Now().Format("/2006-01/"))
	if app.IsExist(saveDir) == false {
		os.MkdirAll(saveDir, os.ModePerm)
	}
	savePath := saveDir + saveFileName
	fW, err := os.Create(savePath)
	if err != nil {
		log.Println("服务端文件创建失败", err)
		resp.Msg = "服务端文件创建失败" + err.Error()
		resp.Code = -1
		return
	}
	log.Println("文件保存路径：", savePath)
	//defer
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		log.Println("服务端文件保存失败", err.Error())
		resp.Msg = "服务端文件保存失败" + err.Error()
		resp.Code = -1
		return
	}
	resp.Data = service.RegVersion(version, project, appName, commitId)

	//resp.Data = uploadToQiniu(savePath, saveFileName)
}

func handlerVersionCheck(writer http.ResponseWriter, request *http.Request) {
	resp := &Resp{Code: 0}
	defer MarshalJson(writer, &resp)
	//
	request.ParseForm()
	form := request.Form
	version := GetParam(form, "version", "").(string)
	commitId := GetParam(form, "commit_id", "").(string)
	appName := GetParam(form, "app", "").(string)
	if version == "" {
		resp.Code = -1
		resp.Data = "参数version不能为空"
		return
	}
	// TODO:判断版本号是否存在 需要提供任务系统接口
	if service.IsExistCommitId(appName, commitId) && commitId != "" {
		resp.Code = -1
		resp.Msg = "该版本已存在发布记录，不允许重复发布，如确实需要重新发布，请勾选重新发布"
		return
	}
	if service.IsExistApp(version, appName) {
		resp.Code = -1
		resp.Msg = "该版本已发布，不允许重复发布，如发布失败重新发布，请勾选重新发布"
		return
	}
	resp.Code = 0
	resp.Msg = "检查通过，允许发布"
}

func handlerVersionReg(writer http.ResponseWriter, request *http.Request) {
	resp := &Resp{Code: 0}
	defer MarshalJson(writer, &resp)
	//
	request.ParseForm()
	form := request.Form
	version := GetParam(form, "version", "").(string)
	project := GetParam(form, "project", "").(string)
	appName := GetParam(form, "app", "").(string)
	commitId := GetParam(form, "commit_id", "").(string)

	if version == "" || project == "" || appName == "" {
		resp.Code = -1
		resp.Data = "必要参数（version，project，app）不能为空"
		return
	}
	// TODO:判断版本号是否存在 需要提供任务系统接口

	resp.Data = service.RegVersion(version, project, appName, commitId)
	resp.Code = 0
	resp.Msg = "注册成功"
}
func handlerUnVersionReg(writer http.ResponseWriter, request *http.Request) {
	resp := &Resp{Code: 0}
	defer MarshalJson(writer, &resp)
	//
	request.ParseForm()
	form := request.Form
	version := GetParam(form, "version", "").(string)
	project := GetParam(form, "project", "").(string)
	appName := GetParam(form, "app", "").(string)
	commitId := GetParam(form, "commit_id", "").(string)

	if version == "" || project == "" || appName == "" {
		resp.Code = -1
		resp.Data = "必要参数（version，project，username）不能为空"
		return
	}
	// TODO:判断任务号是否存在 需要提供任务系统接口

	resp.Data = service.UnRegVersion(version, project, appName, commitId)
	resp.Code = 0
	resp.Msg = "取消注册成功	"
}

func handlerUnknown(writer http.ResponseWriter, request *http.Request) {
	resp := &Resp{Code: 0}
	defer MarshalJson(writer, &resp)
	resp.Code = 0
	resp.Msg = "请求成功"
}

/**
 * 生成一个md5
 */
func CreateMd5(str string) string {
	//sessionId
	m := md5.New()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	io.WriteString(m, str+"_image_crop_"+timestamp+"_"+str)
	return fmt.Sprintf("%x", m.Sum(nil))
}
