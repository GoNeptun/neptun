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

package semafor

import (
	sf "github.com/goneptune/neptune/systemfunctions"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

func Signup(w http.ResponseWriter, r *http.Request) {

	//Check for Maintenance
	var m = sf.CheckMaintenanceMode()

	if !m {

		//Check for session
		var session, _ = sf.CheckSession(r, "")
		var rs = sf.CheckRegistration()
		if !session && rs {
			//Unautorisated person

			switch r.Method {
			case "GET":

				var data = HTMLData{}
				data.HeaderToHTML("Sign Up")  //Title
				data.MenuToHTML(false, false) //Menu
				var bd = []string{"Sign Up Please"}
				data.BodyToHTML(bd) //Content
				data.ShowPage(w, r, "signup.html")

			case "POST":
				// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
				if err := r.ParseForm(); err != nil {
					sf.SetErrorLog(err.Error())
					return
				}
				//fmt.Printf("\n Current timestamp is %s \n", strconv.FormatInt(time.Now().Unix(), 10))

				creds := &SignUpCred{}
				p := bluemonday.UGCPolicy()
				creds.Username = p.Sanitize(r.FormValue("username"))
				creds.Password = p.Sanitize(r.FormValue("password"))
				creds.Mail = p.Sanitize(r.FormValue("mail"))

				//Check mail
				condition := "mail = '" + creds.Mail + "'"
				v := sf.CountRows("*", "users", condition)

				if v == 0 {

					hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
					if err != nil {

						sf.SetErrorLog(err.Error())
					}

					name := []string{"name", "pass", "mail", "created", "status"}
					created := time.Now().Unix()
					value := []string{creds.Username, string(hashedPassword), creds.Mail, strconv.FormatInt(created, 10), "1"}
					sf.InsertRow("users", name, value)

					//check is confirmation email need
					s := sf.IsEmailConfirmation()

					if s != false {
						//Send confirmation email
						SendConfirmationEmail(creds.Mail, creds.Username)
					}
					/////
					var data = HTMLData{}
					data.HeaderToHTML("Sign Up")  //Title
					data.MenuToHTML(false, false) //Menu
					var bd = []string{"You are Signed Up! Please Sign In!"}
					data.BodyToHTML(bd) //Content
					data.ShowPage(w, r, service)

				} else {
					var s = HTMLData{}
					s.MenuToHTML(false, false) //Menu
					answer := "Your Email is not accepted"
					s.ServicePage(answer, w, r)
				}
			}
		} else {
			//Autorisated person
			http.Redirect(w, r, "/", http.StatusFound)
		}

	} else {
		Maintenance(w, r)
	}
}

func SendConfirmationEmail(mail string, uname string) {
	//Send confirmation email
	//1. Create hash
	v := sf.RandomString(64)

	//2. Record hash to db
	//email, hash, timestamp
	nm := []string{"hash", "mail", "name", "created", "param", "deadline"}
	vl := []string{v, mail, uname, strconv.FormatInt(time.Now().Unix(), 10), "confirmemail", strconv.FormatInt(time.Now().Unix(), 10)}
	sf.InsertRow("timehash", nm, vl)

	//3. Send Confirmation email
	//3.1 Create link
	n := sf.LoadConfig("server") // domain is n[2]
	link := "http://" + n[2] + "/user/?param=confirmemail&token=" + v

	//4. Update mailsend timestamp
	lg := strconv.FormatInt(time.Now().Unix(), 10)
	updatedData := [][]string{
		{"mailsent", lg},
	}
	ns := "`name` = '" + uname + "'"
	sf.UpdateRow("users", updatedData, ns)
	/*
			 curldata := [][] string {
				{"to", creds.Mail},
				{"user", creds.Username},
				{"link", link},
			 }
		 go  sf.SysCurl("sendconfmail", curldata)
	*/
	// command go - is run function in background
	go sf.SendEmailNow(mail, uname, link, "Email Confirmation", "first.html") //to, username, message, template
}
