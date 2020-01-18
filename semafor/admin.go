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

//Page for special rights
package semafor

import (
	sf "../system_functions"
	"net/http"
	"strconv"
)

func Admin(w http.ResponseWriter, r *http.Request) {

	//Check for session
	var session, _ = sf.CheckSession(r, "")

	if session != true {
		//Unautorisated person
		http.Redirect(w, r, "/404/", http.StatusFound)

	} else {
		//Autorisated person
		var s = sf.IsSession(r, "admin_token")

		if s != true {
			// No rights for user
			http.Redirect(w, r, "/404/", http.StatusFound)

		} else {

			//update cookie. Need to set time for update. Now it is updated evry time
			sf.UpdateSession(w, r, "session_token")

			var data = HTMLData{}
			data.HeaderToHTML("Admin Panel") //Title
			data.MenuToHTML(true, true)      //Menu
			jsfiles := []string{"admin.js",}
			data.JstoHtml(jsfiles)

//1. Count Users
condition := "`id` > '1'"
userCount := sf.CountRows("*", "users", condition)

//2. Check Site-Settings: Maintenance mode, User Registrations,  Email Confirmations
n := "`id` = '1'"
SqlAnswer := sf.SelectFrom("*", "site_settings", n)
var ss SiteSettingsTable
for SqlAnswer.Next() {

	err := SqlAnswer.StructScan(&ss)
	if err != nil {

		sf.SetErrorLog(err.Error())
	}
}
var Mm string
var MmM string
var MmMc string
if ss.Maintenance == 0 {
	Mm = "Site active"
	MmM = "Switch on Maintenance mode"
	MmMc = "color:green"
} else {
	Mm = "Site in Maintenance"
	MmM = "Switch of Maintenance mode"
	MmMc = "color:red"
}

var Rm string
var RmM string
var Rmc string
if ss.Registration != 0 {
	Rm = "Allow"
	RmM = "Forbid registration"
	Rmc = "color:green"
} else {
	Rm = "Forbiden"
	RmM = "Allow registration"
	Rmc = "color:red"
}

var Mc string
var McM string
var McMc string
if ss.MailConfirmation != 0 {
	Mc = "On"
	McM = "Switch Off Email Confirmation"
	McMc = "color:green"
} else {
	Mc = "Off"
	McM = "Switch On Email Confirmation"
	McMc = "color:red"
}

			var bd = []string{"Admin Panel", strconv.Itoa(userCount), Mm, MmM, MmMc, Rm, RmM, Rmc, Mc, McM, McMc}
			data.BodyToHTML(bd) //Content
			data.ShowPage(w, r, "admin.html")
		}

	}

}
