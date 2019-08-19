package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"encoding/hex"
)

func main() {
	str := "abc123"
	data := []byte(str)
	has := md5.Sum(data)
	fmt.Println(has)
	md5str1 := fmt.Sprintf("%x",has) //将[]byte转成16进制
	fmt.Println(md5str1)

	fmt.Println("----------")
	//方法二
	w:=md5.New()
	io.WriteString(w,str) //将str写入到w中
	bw := w.Sum(nil) //w.Sum(nil)将w的hash转成[]byte格式
	fmt.Println(bw)
	md5str2 := hex.EncodeToString(bw)
	fmt.Println(md5str2)
}
