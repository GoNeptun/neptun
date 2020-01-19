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
	"net/http"
)

func SysAPI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/404/", http.StatusFound)
	case "POST":
		r.ParseForm()
		method := r.Form.Get("method")

		switch method {
		case "sendconfmail":

			SendConfMail(w, r)
		}

	}
}

func SendConfMail(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	to := r.Form.Get("to")
	user := r.Form.Get("user")
	link := r.Form.Get("link")
	sf.SendEmailNow(to, user, link, "Email Confirmation", "first.html")

}
