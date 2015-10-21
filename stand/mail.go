package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Interface interface {
	Alert(int, string) error
}

const (
	ALERT_LEVEL_WARING = iota
	ALERT_LEVEL_MINOR
	ALERT_LEVEL_MAJOR
	ALERT_LEVEL_CRITICAL
)

var (
	User     = "sf_monitor@163.com"
	Password = "cqwsmdgupjmareyc"
	Host     = "smtp.163.com:25"
	ToList   = []string{"me@ckeyer.com", "ckeyer@yeah.net", "cjstudio@yeah.net"}

	alert_title map[int]string
)

func init() {
	title := "监控提示信息"
	alert_title = make(map[int]string)
	alert_title[ALERT_LEVEL_WARING] = `【警告告警】` + title
	alert_title[ALERT_LEVEL_MINOR] = `【次要告警】` + title
	alert_title[ALERT_LEVEL_MAJOR] = `【主要告警】` + title
	alert_title[ALERT_LEVEL_CRITICAL] = `【严重告警】` + title
}

func Alert(lev int, msg string) error {
	switch lev {
	case ALERT_LEVEL_MINOR:
		subject := alert_title[ALERT_LEVEL_MINOR]
		return sendMail(User, Password, Host, ToList, subject, msg)
	case ALERT_LEVEL_MAJOR:
		subject := alert_title[ALERT_LEVEL_MAJOR]
		return sendMail(User, Password, Host, ToList, subject, msg)
	case ALERT_LEVEL_CRITICAL:
		subject := alert_title[ALERT_LEVEL_CRITICAL]
		return sendMail(User, Password, Host, ToList, subject, msg)
	// 默认为警告级别
	default:
		subject := alert_title[ALERT_LEVEL_WARING]
		return sendMail(User, Password, Host, ToList, subject, msg)
	}
}

func sendMail(user, password, host string, send_to []string, subject, body string) error {
	var to string
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	// content_type = "Content-Type: text/html; charset=UTF-8"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	for i, v := range send_to {
		if i == 0 {
			to = v
		} else {
			to += ";" + v
		}
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" +
		content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	// User = "dev@ckeyer.com"
	// Password = "wangcj@Mail0256"
	// Host = "smtp.qq.com:587"

	ToList = []string{"dev@ckeyer.com"}
	err := Alert(ALERT_LEVEL_WARING, "你好，这是一个测试而已")
	if err != nil {
		panic(err)
	}
	fmt.Println("Over...")
}
