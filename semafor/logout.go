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

)

func LogOut(w http.ResponseWriter, r *http.Request) {
	//Check for session
	if !sf.IsSession(r, "") {
		//Unautorisated person - no work

	} else {
		//Autorisated person

		sf.DelSession(w, r, "session_token")
		sf.DelSession(w, r, "admin_token")

	}
	http.Redirect(w, r, "/", http.StatusFound)

}
