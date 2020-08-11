package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

/**
 * 定义Http接口JOSN响应数据结构
 */
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

/**
 * 启动一个Http服务
 */
func HttpService(port string) {

	fmt.Println()
	mux := http.NewServeMux()

	staticFiles := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	//mux.Handle("/", http.StripPrefix("/public/", staticFiles))
	mux.HandleFunc("/upload", handlerUploadRequest)
	mux.HandleFunc("/check", handlerVersionCheck)
	mux.HandleFunc("/reg", handlerVersionReg)
	mux.HandleFunc("/unreg", handlerUnVersionReg)
	mux.HandleFunc("/", handlerUnknown)

	// 启动http服务
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Println("web服务启动,端口号：",port)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("web服务启动失败")
		log.Fatal("listenAndServer : ", err)
	}
	log.Println("web服务启动成功")
}





/**
 * 对象转JSON并响应
 */
func MarshalJson(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, string(data))
	return nil
}

// UnMarshalJson 从request中取出对象
func UnMarshalJson(req *http.Request, v interface{}) error {
	result, err := ioutil.ReadAll(req.Body)
	fmt.Println(req)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(bytes.NewBuffer(result).String()), v)
	return nil
}

func GetParam(params url.Values, name string, defaultValue interface{}) interface{} {
	value, ok := params[name]
	if ok == false {
		return defaultValue
	}
	return value[0]
}
