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
	sf "../system_functions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

func Signin(w http.ResponseWriter, r *http.Request) {

	//Check for session
	var session, _ = sf.CheckSession(r, "")

	if !session {
		//Unautorisated person
		switch r.Method {
		case "GET":

			var data = HTMLData{}
			data.HeaderToHTML("Sign In")  //Title
			data.MenuToHTML(false, false) //Menu
			var bd = []string{"Sign in Please"}
			data.BodyToHTML(bd) //Content
			data.ShowPage(w, r, "signin.html")

		case "POST":

			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			if err := r.ParseForm(); err != nil {
				sf.SetErrorLog(err.Error())
				return
			}
			//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
			creds := &Credentials{}
			creds.Username = r.FormValue("username")
			creds.Password = r.FormValue("password")

			//Check Username
			n := "`name` = '" + creds.Username + "'"
			SqlAnswer := sf.SelectFrom("pass, id, mail_confirmed, login, status, mailsent", "users", n)
			var user UsersTable
			for SqlAnswer.Next() {

				err := SqlAnswer.StructScan(&user)
				if err != nil {

					sf.SetErrorLog(err.Error())
				}
			}
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
				// If the two passwords don't match, return a 401 status
				//w.WriteHeader(http.StatusUnauthorized)
				var s = HTMLData{}
				s.MenuToHTML(false, false) //Menu
				answer := "401"
				s.ServicePage(answer, w, r)

			} else {
				//The Passwords is matched
				var fk bool

				//Check User role for Admin rights
				var m int

				an := "`uid` = '" + user.UID + "'"
				m = sf.CountRows("uid", "roles", an)
				if m != 0 {

					//This is user with special rights

					roleans := sf.SelectFrom("admin", "roles", an)

					var role Roles

					for roleans.Next() {

						err := roleans.StructScan(&role)
						if err != nil {

							sf.SetErrorLog(err.Error())
						}
					}

					//Check if Admin is active
					if role.Admin != "0" {
						//Create Admin Session
						adminsessionToken := sf.SetSession(creds.Username)
						http.SetCookie(w, &http.Cookie{
							Name:    "admin_token",
							Path:    "/",
							Value:   adminsessionToken,
							Expires: time.Now().Add(86400 * time.Second),
						})

					}

				}

				var s bool = sf.CheckMailConfStatus()
				if s {
					//confirmation email is need
					// Check does user confirmed his email
					if user.Mail_Confirmed == 1 {
						fk = true
					} else {
						fk = false
					}

				} else {
					//confirmation email is not need
					fk = true
				}

				lg := strconv.FormatInt(time.Now().Unix(), 10)
				updatedData := [][]string{
					{"access", user.Login},
					{"login", lg},
				}

				sf.UpdateRow("users", updatedData, n)
				//Update new data

				// Create a new random session token
				sessionToken := sf.SetSession(creds.Username)

				// Finally, we set the client cookie for "session_token" as the session token we just generated
				// we also set an expiry time of 120 seconds, the same as the cache
				http.SetCookie(w, &http.Cookie{
					Name:    "session_token",
					Path:    "/",
					Value:   sessionToken,
					Expires: time.Now().Add(86400 * time.Second),
				})

				if fk {

					//get login timestamp


					http.Redirect(w, r, "/", http.StatusFound)

				} else {

					//Check can we resend email
					EmailNotVery(w, r, user.MailSent)
				}
			}
		}
	} else {

		//Autorisated person
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func PreEmailNotVery (w http.ResponseWriter, r *http.Request, uname string ) {
	n := "`name` = '" + uname + "'"
	SqlAnswer := sf.SelectFrom("mailsent", "users", n)
	var user UsersTable
	for SqlAnswer.Next() {

		err := SqlAnswer.StructScan(&user)
		if err != nil {

			sf.SetErrorLog(err.Error())
		}
	}
EmailNotVery(w, r, user.MailSent)

}

func EmailNotVery (w http.ResponseWriter, r *http.Request, mailsend string ) {

	mailsendt, _ := strconv.Atoi(mailsend)
	lg := int(time.Now().Unix())
	tm := lg - mailsendt
	var answer string
	if tm < 900 {
		answer = "E1"
	} else {

		answer = "E2"
	}


	var s = HTMLData{}
	s.MenuToHTML(true, false)
	s.ServicePage(answer, w, r)

}
