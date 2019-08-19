package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var myTemplate *template.Template

type Person struct {
	Name string
	age  string
}
func initTemplate(fileName string) (err error){
	myTemplate,err  = template.ParseFiles(fileName)
	if err != nil{
		fmt.Println("parse file err:",err)
		return
	}
	return
}
func userInfo(w http.ResponseWriter,r *http.Request) {

	p := Person{Name:"Murphy",age:"28"}
	myTemplate.Execute(w,p)

}
func main() {
	initTemplate("./index.html")
	http.HandleFunc("/user/info",userInfo)
	err := http.ListenAndServe("0.0.0.0:8000",nil)
	if err != nil {
		fmt.Println("http listen failed",err)
	}
}