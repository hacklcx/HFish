package send

import (
	"crypto/tls"
	"strconv"
	"gopkg.in/gomail.v2"
	"HFish/utils/log"
)

func SendMail(mailTo []string, subject string, body string, config []string) error {
	port, _ := strconv.Atoi(config[2])
	m := gomail.NewMessage()

	m.SetHeader("From", "<"+config[3]+">")
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(config[0], port, config[3], config[4])
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(m)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "发送邮件通知失败", err)
	} else {
		log.Pr("HFish", "127.0.0.1", "发送邮件通知成功")
	}

	return err
}

func TestMail(addr, protocol, port, account, password string, receivers []string) error {
	intPort, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "<"+account+">")
	m.SetHeader("To", receivers...)    //发送给多个用户
	m.SetHeader("Subject", "HFish测试邮件") //设置邮件主题
	m.SetBody("text/html", "Hello, 这是HFish蜜罐测试邮件！")    //设置邮件正文

	d := gomail.NewDialer(addr, intPort, account, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d.DialAndSend(m)
}