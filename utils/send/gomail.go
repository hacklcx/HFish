package send

import (
	"gopkg.in/gomail.v2"
	"strconv"
	"HFish/utils/log"
	"crypto/tls"
)

func SendMail(mailTo []string, subject string, body string, config []string) error {
	port, _ := strconv.Atoi(config[1])
	m := gomail.NewMessage()

	m.SetHeader("From", "<"+config[2]+">")
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	//d := &gomail.Dialer{
	//	Host:     config[0],
	//	Port:     port,
	//	Username: config[2],
	//	Password: config[3],
	//	SSL:      false,
	//}

	d := gomail.NewDialer(config[0], port, config[2], config[3])

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(m)

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "发送邮件通知失败", err)
	} else {
		log.Pr("HFish", "127.0.0.1", "发送邮件通知成功")
	}

	return err
}
