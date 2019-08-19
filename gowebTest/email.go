package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

/*func SendMail(mailTo []string,subject string, body string ) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string {
		"user": "347443579@qq.com",
		"pass": "mtatamnuiewsbhce",
		"host": "smpt.qq.com",
		"port": "465",
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	//m.SetHeader("From","XD Game" + "<" + mailConn["user"] + ">")  //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("From",mailConn["user"])
	m.SetHeader("To", mailTo...)  //发送给多个用户
	m.SetHeader("Subject", subject)  //设置邮件主题
	m.SetBody("text/html", body)     //设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	if err := d.DialAndSend(m);err != nil{
		return err
	}
	return nil
}
func main()  {
	//定义收件人
	mailTo := []string {
		"1105501072@qq.com",
		"347443579@qq.com",
		"huzhiwei95@gmail.com",
	}
	//邮件主题为"Hello"
	subject := "来啊快活啊"
	// 邮件正文
	body := `<!doctype html>
	<html>
	<head>
	<meta charset="utf-8">
	<title>呵呵</title>
	</head>
	<body>
	fuck
	</body>
	</html>`
	SendMail(mailTo, subject, body)
}*/

func main() {
	m := gomail.NewMessage()

	m.SetAddressHeader("From", "347443579@qq.com" /*"发件人地址"*/, "一位神") // 发件人

	m.SetHeader("To",
		m.FormatAddress("1105501072@qq.com", "收件人")) // 收件人
	/*m.SetHeader("Cc",
		m.FormatAddress("xxxx@foxmail.com", "收件人")) //抄送*/
	/*m.SetHeader("Bcc",
		m.FormatAddress("xxxx@gmail.com", "收件人")) //暗送*/

	m.SetHeader("Subject", "吃屎啦梁非凡")     // 主题

	//m.SetBody("text/html",xxxxx ") // 可以放html..还有其他的
	m.SetBody("text/html","吃屎啦") // 正文

	//m.Attach("我是附件")  //添加附件

	d := gomail.NewDialer("smtp.qq.com", 465, "347443579@qq.com", "mtatamnuiewsbhce") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return
	}

	log.Println("done.发送成功")
}