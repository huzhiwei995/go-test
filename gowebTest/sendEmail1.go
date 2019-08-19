package main

import (
	"net/smtp"
	"strings"
	"fmt"
)

func main() {
	auth := smtp.PlainAuth("","347443579@qq.com","mtatamnuiewsbhce","smtp.qq.com")
	to := []string{"1105501072@qq.com"}
	nickname := "test"
	user := "347443579@qq.com"
	subject := "test mail"
	content_type:= "Content-Type: text/plain; charset=UTF-8"
	body := "吃屎啦梁非凡"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}