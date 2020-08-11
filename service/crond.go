package service

import (
	"log"
	"time"
)

func CrontabFunc(d time.Duration, handler func(), desc string) {
	for {
		log.Println("crontab func runing ", desc)
		handler()
		time.Sleep(d)
	}
}

func Crond() {
	time.Sleep(time.Second * 5)
	go CrontabFunc(time.Second*30, SaveVersionMap, "持久化发布版本数据到本地")
}
