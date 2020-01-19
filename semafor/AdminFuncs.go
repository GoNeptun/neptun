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

//API Page for special rights
//All Data return in JSON

package semafor

import (
	sf "../systemfunctions"
	"net/http"
)

func AdminFuncs(w http.ResponseWriter, r *http.Request) {

	//Check for session
	var session, _ = sf.CheckSession(r, "")

	if session != true {
		//Unautorisated person
		w.Write([]byte(`{"success": 0, "error":"Unautorisated person"}`))

	} else {
		//Autorisated person
		var s = sf.IsSession(r, "admin_token")

		if s != true {
			// No rights for user
			w.Write([]byte(`{"success": 0, "error":"No rights"}`))

		} else {
			switch r.Method {
			case "GET":
				paramm, err := r.URL.Query()["param"]
				param := paramm[0]

				if !err || len(param) < 1 {
					w.Write([]byte(`{"success": 0, "error":"Wrong request"}`))

				} else {

					switch param {
					case "maintenance":
						sf.ParamMaintenance(w, r)
					case "emailcheck":
						sf.EmailCheck(w, r)
					case "regstatus":
						sf.RegStatus(w, r)
					default:
						w.Write([]byte(`{"success": 0, "error":"Wrong request"}`))
					}
				}

			case "POST":
				w.Write([]byte(`{"success": 0, "error":"Wrong method"}`))
			}
		}
	}
}
