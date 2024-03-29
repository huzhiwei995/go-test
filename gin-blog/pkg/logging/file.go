package logging

import (
	"fmt"
	"time"
	"os"
	"log"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s",LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s",LogSaveName,time.Now().Format(TimeFormat),LogFileExt)
	return fmt.Sprint("%s%s",prefixPath,suffixPath)
}
func openLogFile(filePath string) *os.File {
	_,err := os.Stat(filePath)
	switch {
		case os.IsNotExist(err):
			mkDir()
		case os.IsPermission(err):
			log.Fatal("Permission:%v",err)
	}
	handle, err := os.OpenFile(filePath,os.O_APPEND | os.O_CREATE | os.O_WRONLY,0644)
	if err != nil {
		log.Fatalf("fail to openfile:%v",err)
	}
	return handle
}
func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir + "/" + getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}