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
	"fmt"
	"net/http"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {

	//check is user admin
	var s = sf.IsSession(r, "admin_token")
	//Check for Maintenance
	var m = sf.CheckMaintenanceMode()
	if !m || s {
		//Check for session
		var session, uname = sf.CheckSession(r, "")

		if !session {
			//Unautorisated person
			var data = HTMLData{}
			data.HeaderToHTML("Main Page") //Title
			data.MenuToHTML(false, false)  //Menu
			var bd = []string{"Main Page Content"}
			data.BodyToHTML(bd) //Content
			data.ShowPage(w, r, "")
		} else {

			//Autorisated person

			//Check For Email Confirmation and Usres Email Confirmation status
			var m bool = sf.UserMailConf(uname)
			if m {
				//User confirm email or email confirmation disabled

				//update cookie. Need to set time for update. Now it is updated evry time
				sf.UpdateSession(w, r, "session_token")

				var data = HTMLData{}
				data.HeaderToHTML("Dashboard")                        //Title
				data.MenuToHTML(true, sf.IsSession(r, "admin_token")) //Menu
				var bd = []string{fmt.Sprintf("Dashboard of %s", uname)}
				data.BodyToHTML(bd) //Content
				data.ShowPage(w, r, "")

			} else {
				//User  don't confirm email
				PreEmailNotVery(w, r, uname)
			}

			//end ayth user
		}

	} else {
		Maintenance(w, r)
	}

}
