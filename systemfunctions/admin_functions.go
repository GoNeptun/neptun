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

package systemfunctions

import (
	"net/http"
)

func ParamMaintenance(w http.ResponseWriter, r *http.Request) {
	//Switch On / Off Maintenance Mode
	var m bool = CheckMaintenanceMode()
	if !m {
		// put to maintenance
		ToMaintenance()
		//Send answer
		w.Write([]byte(`{"success": 1, "answer":{"status": "In Maintenance", "style":"red", "button":"Switch off Maintenance mode"}}`))
	} else {
		//make alive
		MakeAlive()
		//Send answer
		w.Write([]byte(`{"success": 1, "answer":{"status": "Site active", "style":"green", "button":"Switch on Maintenance mode"}}`))
	}
}

func RegStatus(w http.ResponseWriter, r *http.Request) {
	//Switch On / Off user registration
	var m bool = CheckRegistration()
	if m {
		//Forbid registration
		ForbidRegistration()
		w.Write([]byte(`{"success": 1, "answer":{"status": "Forbidden", "style":"red", "button":"Allow User Registration"}}`))
	} else {
		//Allow registration
		AllowRegistration()
		w.Write([]byte(`{"success": 1, "answer":{"status": "Allow", "style":"green", "button":"Forbid User Registration"}}`))
	}
}

func EmailCheck(w http.ResponseWriter, r *http.Request) {
	//Switch On / Off email confirmatio
	var m bool = CheckMailConfStatus()

	if m {
		//Switch  Off email confirmation
		OffMailConfirmation()
		w.Write([]byte(`{"success": 1, "answer":{"status": "Off", "style":"red", "button":"Switch on Email Confirmation"}}`))
	} else {
		//Switch  On email confirmation
		OnMailConfirmation()
		w.Write([]byte(`{"success": 1, "answer":{"status": "On", "style":"green", "button":"Switch off Email Confirmation"}}`))
	}
}
