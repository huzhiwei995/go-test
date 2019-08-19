package main

import (
	"encoding/json"
	"fmt"
)

type name map[string]struct{
	Name string `json:"name"`
	Age int `json:"age"`
}
type names struct{
	Name string `json:"name"`
	Age int `json:"age"`
}
func main(){
	str := `{"1":{"name":"asdas","age":10}}`
	var data name
	err := json.Unmarshal([]byte(str),&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
/*func main(){
	str := `{"name":"asdas","age":10}`
	var data names
	err := json.Unmarshal([]byte(str),&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}*/