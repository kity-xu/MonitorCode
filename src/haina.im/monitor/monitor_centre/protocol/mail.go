package protocol

import (
	"github.com/cihub/seelog"
	"net/smtp"
	"strings"
)

/*
 * user : example@example.com login smtp server user
 * password: xxxxx login smtp server password
 * host: smtp.example.com:port   smtp.163.com:25
 * to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */

type Notification struct {
}

func configEmail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
func (this *Notification) SendMail(message string) {
	log, _ := seelog.LoggerFromConfigAsFile("haina.im/monitor/monitor_centre/config/logconfig.xml")
	defer log.Flush()
	user := "cherish_xulang@163.com"
	password := "xxdlyn1025"
	host := "smtp.163.com:25"
	to := "474493616@qq.com;2865794478@qq.com"
	subject := "Test send email by monitor_node"

	body := `
	<html>
 	<body>
 	<h3>` + message + `</h3>
 	</body>
 	</html>
 	`
	log.Debug("send email")
	err := configEmail(user, password, host, to, subject, body, "html")
	if err != nil {
		log.Debug("send mail error!")
		log.Debug(err)
	} else {
		log.Debug("send mail success!")
	}
}

func (this *Notification) SendSms(message string) {

}
