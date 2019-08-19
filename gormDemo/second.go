package gormDemo

import "fmt"

/*type Student struct {
	Student
}*/
func main()  {
	var stu Student
	stu.Age = 12
	stu.Name="胡志伟"
	fmt.Println(stu)
}
