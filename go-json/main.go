package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func main() {
	file,err := ioutil.ReadFile("./json.txt")
	if err != nil {
		fmt.Println("error : ",err)
	}
	str := string(file)
	var v map[string]interface{}
	err = json.Unmarshal([]byte(str),&v)
	if err != nil {
		fmt.Println("error : ",err)
	}
	fmt.Println(v)
}
