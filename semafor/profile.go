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
	"net/http"
	"strconv"
	"time"
)

func Profile(w http.ResponseWriter, r *http.Request) {

	//Check for session
	var session, uname = sf.CheckSession(r, "")

	if session != true {
		//Unautorisated person
		http.Redirect(w, r, "/", http.StatusFound)

	} else {
		//Autorisated person

		//update cookie. Need to set time for update. Now it is updated evry time
		sf.UpdateSession(w, r, "session_token")

		var data = HTMLData{}
		data.HeaderToHTML("Profile")                          //Title
		data.MenuToHTML(true, sf.IsSession(r, "admin_token")) //Menu

		//Create Content
		var a UsersTable
		n := "`name` = '" + uname + "'"
		SqlAnswer := sf.SelectFrom("mail, created, login", "users", n)
		for SqlAnswer.Next() {

			err := SqlAnswer.StructScan(&a)
			if err != nil {

				sf.SetErrorLog(err.Error())
			}
		}

		jointime, _ := strconv.ParseInt(a.Created, 10, 64)
		logintime, _ := strconv.ParseInt(a.Login, 10, 64)
		jt := time.Unix(jointime, 0).Format("02 January 2006")
		lt := time.Unix(logintime, 0).Format("02 January 2006 15:04:05")

		var bd = []string{"Profile", uname, string(jt), string(lt), a.Mail}
		data.BodyToHTML(bd) //Content
		data.ShowPage(w, r, "profile.html")

	}

}
