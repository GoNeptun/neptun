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
	"html/template"
	"net/http"
)

func (s *HTMLData) ServicePage(name string, w http.ResponseWriter, r *http.Request) {

	if name == "404" {
		s.HeaderToHTML("Error 404")
		var bd = []template.HTML{"Error 404. Page Not Found"}
		s.HTMLBodyToHTML(bd)
	}

	if name == "401" {
		s.HeaderToHTML("Error 401")
		var bd = []string{"Email or Password is wrong"}
		s.BodyToHTML(bd)
	}

	if name == "500" {
		s.HeaderToHTML("Error 500")
		var bd = []string{"Error 500. Internal Server Error"}
		s.BodyToHTML(bd)
	}

	if name == "Wrong Link" {
		s.HeaderToHTML("Wrong Link")
		var bd = []string{"Email Not confirmed. Wrong Link"}
		s.BodyToHTML(bd)
	}

	if name == "Email confirmed" {
		s.HeaderToHTML(name)
		var bd = []string{name}
		s.BodyToHTML(bd)
	}

	if name == "Email not confirmed" {
		s.HeaderToHTML(name)
		var bd = []string{name}
		s.BodyToHTML(bd)
	}

	if name == "E3" {
		s.HeaderToHTML("Hash time if over")
		var bd = []template.HTML{"Hash time if over. Please request a new link"}
		s.HTMLBodyToHTML(bd)
	}

	if name == "E1" {
		//Email not confirmed and confirmation link was sent less
		//then 15 minutes ago
		s.HeaderToHTML("Email not confirmed")
		var bd = []template.HTML{"Email not confirmed.", " Please check your email for confirmation link. You can request new confirmation email in 15 minutes"}
		s.HTMLBodyToHTML(bd)
	}

	if name == "A1" {
		//Email not confirmed and confirmation link was sent less
		//then 15 minutes ago
		s.HeaderToHTML("Email successfully changed")
		var bd = []template.HTML{"Email successfully changed."}
		s.HTMLBodyToHTML(bd)
	}
	if name == "E2" {
		//Email not confirmed and confirmation link was sent less
		//then 15 minutes ago
		s.HeaderToHTML("Email not confirmed")
		var bd = []template.HTML{"Email not confirmed.", "Please check your email for confirmation link or <span id='reqnewemail' style='cursor:pointer; text-decoration: underline;'>request a new email</span>"}
		s.HTMLBodyToHTML(bd)
		jsfiles := []string{"query.js"}
		s.JstoHtml(jsfiles)
	}

	if name == "Your Email is not accepted" {
		s.HeaderToHTML(name)
		var bd = []string{name}
		s.BodyToHTML(bd)
	}

	if name == "Password Changed" || name == "The passwords are not the same" || name == "There is an error while change password" {
		s.HeaderToHTML(name)
		var bd = []string{name}
		s.BodyToHTML(bd)
	}
	s.ShowPage(w, r, "service")

}

func NotFoundAny(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/404/", http.StatusFound)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	//Check for special page Request

	//check is user admin
	var s = sf.IsSession(r, "admin_token")
	//Check for Maintenance
	var m = sf.CheckMaintenanceMode()
	if !m || s {

		var session, _ = sf.CheckSession(r, "")
		var lk = HTMLData{}
		lk.MenuToHTML(true, sf.IsSession(r, "admin_token"))

		if !session {
			//Unautorisated person

			lk.MenuToHTML(false, false)
		}

		lk.ServicePage("404", w, r)

	} else {
		Maintenance(w, r)
	}

}
