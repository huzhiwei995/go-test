package main

import (
	"net/http"
	"log"
	"os"
	"io/ioutil"
	"io"
)

var (
	Trace *log.Logger
	Info *log.Logger
	Warning  *log.Logger
	Error *log.Logger
)

func init(){
	file,err := os.OpenFile("errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666,
	)
	if err != nil {
		log.Fatalln("Fatled to open error log file:",err)
	}
	Trace = log.New(ioutil.Discard,"TRACE: ",log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout,"INFO",log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout,"WARNING",log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file,os.Stderr),"ERROR: ",log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	url := "http://www.baidu.com"
	httpsrc,err := http.Get(url)
	if err != nil {
		Trace.Println(err)
	}
	defer httpsrc.Body.Close()

}
