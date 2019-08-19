package main

import (
	"encoding/json"
	"log"
	"fmt"
)

var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"types":1,
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	//json处理
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	for _,v := range c {
		fmt.Println(v)
		fmt.Printf("%T\n",v)
		fmt.Println("--------")
	}
	fmt.Println("============")
	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact")
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])
}