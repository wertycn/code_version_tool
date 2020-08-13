package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

func GetMd5(str string) string {
	m := md5.New()
	//timestamp := time.Now().Format("2006-01-02 15:04:05")
	io.WriteString(m, "f10_"+str)
	return fmt.Sprintf("%x", m.Sum(nil))
}


func StrTime2Int(toBeCharge string) int64 {
	//toBeCharge := "2015-01-01 00:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()
	return sr
}