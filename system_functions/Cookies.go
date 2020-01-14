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

package system_functions

import (
	"net/http"
	"time"
)

// addCookie will apply a new cookie to the response of a http
// request, with the key/value this method is passed.
func addCookie(w http.ResponseWriter, name string, value string) {
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

// Coockies set:
//addCookie(w, "TestCookieName", "TestValue")
