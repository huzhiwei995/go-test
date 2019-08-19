package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"ginDemo"
)
type Product struct {
	gorm.Model
	Code string
	Price uint
}

type Model struct {
	ginDemo.Student
}
/*func main() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
}*/

func main() {
	var stu Model
	stu.Age = 18
	stu.Name = "胡志伟"
	fmt.Println(stu)
}