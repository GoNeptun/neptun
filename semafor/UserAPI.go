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
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

func UserAPI(w http.ResponseWriter, r *http.Request) {

	//check is user admin
	var s = sf.IsSession(r, "admin_token")
	//Check for Maintenance
	var m = sf.CheckMaintenanceMode()
	if !m || s {

		var session, _ = sf.CheckSession(r, "")
		var s = HTMLData{}
		s.MenuToHTML(true, sf.IsSession(r, "admin_token")) //Menu

		if !session {
			//Unautorisated person
			s.MenuToHTML(false, false) //Menu
		}

		switch r.Method {
		case "GET":

			param, _ := r.URL.Query()["param"]
			p := bluemonday.UGCPolicy()
			paramm := p.Sanitize(param[0])
			switch paramm {
			case "confirmemail":
				ConfirmEmail(w, r)
			case "forgot":
				//User ask forgot from his email link
				FP(w, r)
			case "reqnewemail":
				ReqNewEmail(w, r)
			case "changeemail":
				ChangeUserEmail(w, r)
			case "changepass":
				ChangePass(w, r)
			default:
				NotFoundAny(w, r)
			}

		case "POST":
			err := r.ParseForm()
			if err != nil {
				sf.SetErrorLog(err.Error())
			}
			v := r.Form
			param := v.Get("param")
			p := bluemonday.UGCPolicy()
			paramm := p.Sanitize(param)
			switch paramm {
			case "restorepass":
				FP(w, r)
			default:
				NotFoundAny(w, r)
			}

		}
	} else {
		Maintenance(w, r)
	}

}

func ChangePass(w http.ResponseWriter, r *http.Request) {
	//1. Check IsSession
	var session, uname = sf.CheckSession(r, "")

	if session {

		// Check User Password
		n := "`name` = '" + uname + "'"
		SqlAnswer := sf.SelectFrom("pass", "users", n)
		var user UsersTable
		for SqlAnswer.Next() {

			err := SqlAnswer.StructScan(&user)
			if err != nil {

				sf.SetErrorLog(err.Error())
			}
		}
		p := bluemonday.UGCPolicy()
		pas, err := r.URL.Query()["oldpass"]
		if err {
			sf.SetErrorLog("No Get['oldpass']")
		}
		pass := p.Sanitize(pas[0])
		if !err || len(pass) > 1 {

			newpas, err := r.URL.Query()["newpass"]
			if err {
				sf.SetErrorLog("No Get['newpass']")
			}
			newpass := p.Sanitize(newpas[0])
			if !err || len(newpass) > 1 {

				if newpass != pass {

					//Check password
					if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
						// If the two passwords don't match
						w.Write([]byte(`{"success": 0, "error":"Wrong Password"}`))
					} else {

						//Rewrite password

						hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newpass), 8)
						if err != nil {

							sf.SetErrorLog(err.Error())
						}
						stringpass := "'" + string(hashedPassword) + "'"
						updatedData := [][]string{
							{"`pass`", stringpass},
						}
						nv := "`name` = '" + uname + "'"
						sf.UpdateRow("users", updatedData, nv)

						w.Write([]byte(`{"success": 1, "answer":"The Password successfully updated"}`))

					}
				} else {
					w.Write([]byte(`{"success": 0, "error":"Your New Password is same with New one"}`))
				}
			} else {
				w.Write([]byte(`{"success": 0, "error":"Wrong Request"}`))
			}

		} else {
			w.Write([]byte(`{"success": 0, "error":"Wrong Request"}`))
		}

	} else {
		w.Write([]byte(`{"success": 0, "error":"Unautorisated person"}`))
	}
}

func ChangeUserEmail(w http.ResponseWriter, r *http.Request) {
	//1. Check IsSession
	var session, uname = sf.CheckSession(r, "")

	if session {

		// Check User Password and email
		n := "`name` = '" + uname + "'"
		SqlAnswer := sf.SelectFrom("pass, mail", "users", n)
		var user UsersTable
		for SqlAnswer.Next() {

			err := SqlAnswer.StructScan(&user)
			if err != nil {

				sf.SetErrorLog(err.Error())
			}
		}
		p := bluemonday.UGCPolicy()
		pas, err := r.URL.Query()["pass"]
		if err {
			sf.SetErrorLog("No Get['pass']")
		}
		pass := p.Sanitize(pas[0])
		if !err || len(pass) > 1 {
			mail, err := r.URL.Query()["email"]
			if err {
				sf.SetErrorLog("No Get['email']")
			}
			email := p.Sanitize(mail[0])
			if !err || len(email) > 1 {

				if email != user.Mail {
					//Check password
					if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
						// If the two passwords don't match
						w.Write([]byte(`{"success": 0, "error":"Wrong Password"}`))
					} else {
						// Check does the email is unical
						condition := "mail = '" + email + "'"
						v := sf.CountRows("*", "users", condition)

						if v == 0 {
							//Send new confirmation email

							SendConfirmationEmail(email, uname)

							w.Write([]byte(`{"success": 1, "answer":"Please check your new email for confirmation letter"}`))

						} else {
							w.Write([]byte(`{"success": 0, "error":"Email is same"}`))
						}
					}
				} else {
					w.Write([]byte(`{"success": 0, "error":"Email is same"}`))
				}

			} else {
				w.Write([]byte(`{"success": 0, "error":"Wrong Request"}`))
			}

		} else {
			w.Write([]byte(`{"success": 0, "error":"Wrong Request"}`))
		}

	} else {
		w.Write([]byte(`{"success": 0, "error":"Unautorisated person"}`))
	}
}

func ReqNewEmail(w http.ResponseWriter, r *http.Request) {
	//1. Check IsSession
	var session, uname = sf.CheckSession(r, "")

	if session {
		//2. Check users
		n := "`name` = '" + uname + "'"
		SqlAnswer := sf.SelectFrom("name, mail, mail_confirmed, mailsent", "users", n)
		var user UsersTable
		for SqlAnswer.Next() {

			err := SqlAnswer.StructScan(&user)
			if err != nil {

				sf.SetErrorLog(err.Error())
			}
		}
		if user.MailConfirmed == 1 {
			w.Write([]byte(`{"success": 0, "error":"Error request"}`))
		} else {

			mailsendt, _ := strconv.Atoi(user.MailSent)
			lg := int(time.Now().Unix())
			tm := lg - mailsendt

			if tm < 900 {
				w.Write([]byte(`{"success": 0, "error":"Error request"}`))
			} else {
				//Delete all hash data about this mail
				condition := "mail = '" + user.Mail + "'"
				vm := sf.CountRows("*", "timehash", condition)
				if vm != 0 {
					sf.DeleteRow("timehash", condition)
				}
				//4. Send email
				SendConfirmationEmail(user.Mail, user.Name)
				w.Write([]byte(`{"success": 1, "answer":"The Confirmation email was sent"}`))
			}
		}
	} else {
		w.Write([]byte(`{"success": 0, "error":"Unautorisated person"}`))
	}
}

func FP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//Check the token and show page for change password
		var answer string
		p := bluemonday.UGCPolicy()
		hash, err := r.URL.Query()["token"]
		if err {
			sf.SetErrorLog("No Get['token']")
		}
		hashm := p.Sanitize(hash[0])
		if !err || len(hashm) > 1 {

			//Check hash
			condition := "hash = '" + hashm + "'"
			v := sf.CountRows("*", "timehash", condition)
			if v != 0 {

				//Check hashtime
				ans := sf.SelectFrom("deadline", "timehash", condition)
				var u TimeHash
				for ans.Next() {

					err := ans.StructScan(&u)
					if err != nil {

						sf.SetErrorLog(err.Error())
					}
				}

				nowtime := time.Now().Unix()
				if u.Deadline > int(nowtime) {
					//show page
					var data = HTMLData{}
					data.HeaderToHTML("Create new Password") //Title
					data.MenuToHTML(false, false)            //Menu
					var bd = []string{"Create new Password", hashm}
					data.BodyToHTML(bd) //Content
					data.ShowPage(w, r, "restore_pass.html")
				} else {
					var s = HTMLData{}
					s.MenuToHTML(false, false) //Menu
					answer = "E3"
					s.ServicePage(answer, w, r)
				}

			} else {
				var s = HTMLData{}
				s.MenuToHTML(false, false) //Menu
				answer = "Wrong Link"
				s.ServicePage(answer, w, r)
			}
		} else {
			var s = HTMLData{}
			s.MenuToHTML(false, false) //Menu
			answer = "Wrong Link"
			s.ServicePage(answer, w, r)
		}

	case "POST":
		//Get data and change password
		v := r.Form
		p := bluemonday.UGCPolicy()
		passa := p.Sanitize(v.Get("password_a"))
		passb := p.Sanitize(v.Get("password_b"))
		token := p.Sanitize(v.Get("token"))

		//Check hash
		condition := "hash = '" + token + "'"
		vm := sf.CountRows("*", "timehash", condition)
		if vm != 0 {

			if passa == passb {
				//Get user email
				n := "`hash` = '" + token + "'"
				SqlAnswer := sf.SelectFrom("mail", "timehash", n)
				var user TimeHash
				for SqlAnswer.Next() {

					err := SqlAnswer.StructScan(&user)
					if err != nil {

						sf.SetErrorLog(err.Error())
					}
				}
				//Update the passwords
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passa), 8)
				if err != nil {

					sf.SetErrorLog(err.Error())
				}
				stringpass := "'" + string(hashedPassword) + "'"
				updatedData := [][]string{
					{"`pass`", stringpass},
				}
				nv := "`mail` = '" + user.Mail + "'"
				sf.UpdateRow("users", updatedData, nv)
				//Delete hash
				cond := "`hash` = '" + token + "'"
				sf.DeleteRow("timehash", cond)
				//Return page
				var s = HTMLData{}
				s.MenuToHTML(false, false)
				answer := "Password Changed"
				s.ServicePage(answer, w, r)

			} else {
				//The passwords are not the same
				var s = HTMLData{}
				s.MenuToHTML(false, false)
				answer := "The passwords are not the same"
				s.ServicePage(answer, w, r)
			}
		} else {
			//No token
			var s = HTMLData{}
			s.MenuToHTML(false, false)
			answer := "There is an error while change password"
			s.ServicePage(answer, w, r)
		}
	}
}

func ConfirmEmail(w http.ResponseWriter, r *http.Request) {

	//Check link
	var answer string
	p := bluemonday.UGCPolicy()
	hash, err := r.URL.Query()["token"]
	if err {
		sf.SetErrorLog("No GET['token']")
	}
	hashm := p.Sanitize(hash[0])

	if !err || len(hashm) > 1 {

		//Check hash
		condition := "hash = '" + hashm + "'"
		v := sf.CountRows("*", "timehash", condition)

		if v != 0 {

			//Check hashtime
			ans := sf.SelectFrom("deadline", "timehash", condition)
			var u TimeHash
			for ans.Next() {

				err := ans.StructScan(&u)
				if err != nil {

					sf.SetErrorLog(err.Error())
				}
			}

			nowtime := time.Now().Unix()

			if u.Deadline > int(nowtime) {

				//1. Get email daress from DB
				sqldata := sf.SelectFrom("mail, name", "timehash", condition)
				var m TimeHash
				for sqldata.Next() {

					err := sqldata.StructScan(&m)
					if err != nil {

						sf.SetErrorLog(err.Error())
					}
				}
				//mail = m.Mail

				//1.2 //Check if email is mismath
				cv := "`name` = '" + m.Name + "'"
				sq := sf.SelectFrom("mail, name", "users", cv)
				var u UsersTable
				for sq.Next() {

					err := sq.StructScan(&u)
					if err != nil {

						sf.SetErrorLog(err.Error())
					}
				}

				if u.Mail != m.Mail {
					// User want to change his email
					//2. Update that email confirmed in table users
					cond := "`name` = '" + m.Name + "'"
					updatedData := [][]string{
						{"`mail_confirmed`", "'1'"},
						{"`mail`", "'" + m.Mail + "'"},
					}

					sf.UpdateRow("users", updatedData, cond)

					sf.DeleteRow("timehash", cond)

					//4. Return sucsess data
					answer = "A1"

				} else {

					//2. Update that email confirmed in table users
					cond := "mail = '" + m.Mail + "'"
					updatedData := [][]string{
						{"mail_confirmed", "1"},
					}

					sf.UpdateRow("users", updatedData, cond)
					//3. Delete hash from timehash table
					sf.DeleteRow("timehash", condition)
					//4. Return sucsess data
					answer = "Email confirmed"
				}
			} else {
				answer = "E3"
			}
		} else {
			answer = "Wrong Link"
		}

	} else {
		answer = "Wrong Link"
	}

	var s = HTMLData{}
	s.MenuToHTML(sf.IsSession(r, ""), sf.IsSession(r, "admin_token")) //Menu
	s.ServicePage(answer, w, r)
}
