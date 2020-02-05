// Copyright 2019 Alexey Yanchenko <mail@yanchenko.me>
//
// This file is part of the Neptune library.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package systemfunctions

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"path"
)

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {

	n := LoadConfig("email")
	//n[0] = host, n[1] = port, n[2] = user, n[3] = password, n[4] = address, n[5] = reply to

	var auth smtp.Auth
	// PlainAuth(identity, username, password, host string)
	// Usually identity should be the empty string, to act as username.

	auth = smtp.PlainAuth("", n[4], n[3], n[0])

	fromheader := "From: Support <" + n[4] + ">\r\n"
	toheader := "Reply-To: " + n[5] + "\n"
	xm := "X-Mailer: php \n"
	mime := fromheader + toheader + xm + "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := n[0] + ":" + n[1]

	//SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	//The msg headers should usually include fields such as "From", "To", "Subject", and "Cc".

	if err := smtp.SendMail(addr, auth, n[4], r.to, msg); err != nil {
		return false, err
		SetErrorLog(err.Error())
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err

	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err

	}
	r.body = buf.String()
	return nil
}

func SendEmailNow(to string, uname string, link string, subject string, template string) {

	templateData := struct {
		Name string
		URL  string
	}{
		Name: uname,
		URL:  link,
	}
	n := LoadContentDirectory()
	templateDir := n + "layouts/template/email/"
	var emailtemplate = path.Join(templateDir, template)
	r := NewRequest([]string{to}, subject, "")

	if err := r.ParseTemplate(emailtemplate, templateData); err == nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)

	} else {
		SetErrorLog(err.Error())
	}
}
