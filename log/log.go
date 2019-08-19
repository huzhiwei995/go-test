package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	logName := "./test.log"
	Debug(logName)
}
func Debug(logName string) {
	logFile,err := os.OpenFile(logName,os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("create ./test.log err : %v\n", err)
	}
	defer logFile.Close()
	debugLog := log.New(logFile,"[DEBUG]",log.Ldate)
	debugLog.SetPrefix("[Debug]")
	debugLog.SetFlags(log.Lshortfile)
	debugLog.Println("this is Debug log")
}