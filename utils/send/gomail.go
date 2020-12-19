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
	m.SetHeader("To", mailTo...)    //Send to multiple users
	m.SetHeader("Subject", subject) //Set email subject
	m.SetBody("text/html", body)    //Set the message body

	d := gomail.NewDialer(config[0], port, config[3], config[4])
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(m)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Failed to send email notification", err)
	} else {
		log.Pr("HFish", "127.0.0.1", "Send email notification success")
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
	m.SetHeader("To", receivers...)    //Send to multiple users
	m.SetHeader("Subject", "HFish test mail") //Set email subject
	m.SetBody("text/html", "Hello, This is the HFish honeypot test email！")    //Set the message body

	d := gomail.NewDialer(addr, intPort, account, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d.DialAndSend(m)
}
