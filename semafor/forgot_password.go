// Copyright 2019 Alexey Yanchenko <mail@yanchenko.me>
//
// This file is part of the Neptun library.
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
	sf "../systemfunctions"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"strconv"
	"time"
)

func ForgotPassword(w http.ResponseWriter, r *http.Request) {

	//Check for Maintenance
	var m = sf.CheckMaintenanceMode()
	if !m {
		//Check for session
		var session, _ = sf.CheckSession(r, "")

		if !session {
			//Unautorisated person
			switch r.Method {
			case "GET":
				var bd = []string{"Forgot Password"}
				ReturnFPPage(w, r, "Forgot Password", bd)

			case "POST":

				// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
				if err := r.ParseForm(); err != nil {
					sf.SetErrorLog(err.Error())
					return
				}
				//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
				p := bluemonday.UGCPolicy()
				creds := &ForgotPass{}
				creds.Mail = p.Sanitize(r.FormValue("email"))

				//Check mail
				condition := "mail = '" + creds.Mail + "'"
				v := sf.CountRows("*", "users", condition)

				if v == 0 {
					//There is no this email or it is not email
				} else {
					//this email is in system
					//Send confirmation email
					//1. Create hash
					v := sf.RandomString(64)

					//2. Record hash to db
					//email, hash, timestamp
					nm := []string{"hash", "mail", "created", "param", "deadline"}
					deadline := time.Now().Unix() + 172800
					vl := []string{v, creds.Mail, strconv.FormatInt(time.Now().Unix(), 10), "forgot", strconv.FormatInt(deadline, 10)}
					sf.InsertRow("timehash", nm, vl)

					//3. Send Confirmation email
					//3.1 Create link
					n := sf.LoadConfig("server") // domain is n[2]
					link := "http://" + n[2] + "/user/?param=forgot&token=" + v
					go sf.SendEmailNow(creds.Mail, creds.Mail, link, "Reset the Password", "restore_pass_mail.html") //to, username, message, template
				}
				var bd = []string{"Please check your email for restore password code"}
				ReturnFPPage(w, r, "Forgot Password", bd)

			}
		} else {
			//Autorisated person
			http.Redirect(w, r, "/", http.StatusFound)
		}

	} else {
		Maintenance(w, r)
	}

}

func ReturnFPPage(w http.ResponseWriter, r *http.Request, title string, content []string) {
	var data = HTMLData{}
	data.HeaderToHTML(title)      //Title
	data.MenuToHTML(false, false) //Menu
	data.BodyToHTML(content)      //Content
	data.ShowPage(w, r, "forgot.html")

}
