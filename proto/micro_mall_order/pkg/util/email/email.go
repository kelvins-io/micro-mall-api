package email

import (
	"context"
	"crypto/tls"
	"fmt"
	"gitee.com/kelvins-io/kelvins"
	"net/smtp"
	"strings"
)

type SendRequest struct {
	Receivers []string `json:"receivers"`
	Subject   string   `json:"subject"`
	Message   string   `json:"message"`
}

type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type Client struct {
	config *Config
}

func NewClient(user, pwd, host, port string) *Client {
	return &Client{config: &Config{
		User:     user,
		Password: pwd,
		Host:     host,
		Port:     port,
	}}
}

func (e *Client) buildMessage(req *SendRequest) string {
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("From: 重要提醒<%s>\r\n", e.config.User))
	if len(req.Receivers) > 0 {
		buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(req.Receivers, ";")))
	}
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", req.Subject))
	buf.WriteString("Content-Type: text/html; charset=UTF-8")
	buf.WriteString("\r\n\r\n")
	buf.WriteString(req.Message)

	return buf.String()
}

func (e *Client) SendEmail(req *SendRequest) error {

	messageBody := e.buildMessage(req)
	//build an auth
	auth := smtp.PlainAuth("", e.config.User, e.config.Password, e.config.Host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         e.config.Host,
	}
	serverName := e.config.Host + ":" + e.config.Port

	conn, err := tls.Dial("tcp", serverName, tlsConfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, e.config.Host)
	if err != nil {
		return err
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		return err
	}

	// step 2: add all from and to
	if err = client.Mail(e.config.User); err != nil {
		return err
	}
	for _, k := range req.Receivers {
		if err = client.Rcpt(k); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	err = client.Quit()
	if err != nil {
		kelvins.ErrLogger.Errorf(context.Background(), "client quit err: %v", err)
		return err
	}

	return nil
}
