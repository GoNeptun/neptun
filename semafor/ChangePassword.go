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
	sf "../systemfunctions"
	"net/http"
)

func ChangePassword(w http.ResponseWriter, r *http.Request) {

	//Check for session
	var session, _ = sf.CheckSession(r, "")

	if session != true {
		//Unautorisated person
		http.Redirect(w, r, "/", http.StatusFound)

	} else {
		//Autorisated person

		//update cookie. Need to set time for update. Now it is updated evry time
		sf.UpdateSession(w, r, "session_token")

		var data = HTMLData{}
		data.HeaderToHTML("Change Password")                  //Title
		data.MenuToHTML(true, sf.IsSession(r, "admin_token")) //Menu
		jsfiles := []string{"query.js"}
		data.JstoHtml(jsfiles)

		var bd = []string{"Change Password"}
		data.BodyToHTML(bd) //Content
		data.ShowPage(w, r, "change_password.html")

	}

}
