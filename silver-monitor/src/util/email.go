package util

import (
    "net/smtp"
    "strings"
    "bytes"
    "fmt"
    "time"
)

// 发送邮件
func EmailSend(user string, password string, host string, to string, subject string, body string) error {
    hp := strings.Split(host, ":")
    auth := smtp.PlainAuth("", user, password, hp[0])
    send_to := strings.Split(to, ";")

    buffer := bytes.NewBuffer(nil)

    boudary := "SILVER_MONITOR"

    msg := fmt.Sprintf("To:%s\r\n" +
    "From:%s\r\n" +
    "Subject:%s\r\n" +
    "Content-Type:multipart/mixed;Boundary=\"%s\"\r\n" +
    "Mime-Version:1.0\r\n" +
    "Date:%s\r\n" +
    "\r\n\r\n--%s\r\n" +
    "Content-Type:text/plain;charset=utf-8\r\n\r\n%s\r\n",
        to, user, subject, boudary, time.Now().String(), boudary, body)

    buffer.WriteString(msg)

    err := smtp.SendMail(host, auth, user, send_to, buffer.Bytes())

    return err
}
