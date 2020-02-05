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
	"fmt"
	"net/http"
)

func CheckSession(r *http.Request, sesname string) (m bool, v string) {
	if sesname != "" {

	} else {
		sesname = "session_token"
	}
	c, err := r.Cookie(sesname)
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status

			m = false
			v = ""
			return m, v
		}
		// For any other type of error, return a bad request status

		m = false
		v = ""
		return m, v

	}
	sessionToken := c.Value

	// We then get the name of the user from our cache, where we set the session token
	response, err := cache.Do("GET", sessionToken)
	if err != nil {
		// If there is an error fetching from cache, return an internal server error status

		m = false
		v = ""
		return m, v

	}

	if response == nil {
		// If the session token is not present in cache, return an unauthorized error

		m = false
		v = ""
		return m, v

	}
	//Check is ok
	m = true
	v = fmt.Sprintf("%s", response)

	return m, v

}

func IsSession(r *http.Request, sesname string) bool {
	var session, _ = CheckSession(r, sesname)
	return session
}
