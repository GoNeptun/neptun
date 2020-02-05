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
	"github.com/gomodule/redigo/redis"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

func SetSession(name string) string {
	// Create a new random session token

	sessionToken := uuid.Must(uuid.NewV4()).String()
	sessionToken = string(sessionToken)
	// Set the token in the cache, along with the user whom it represents
	// The token has an expiry time of 120 seconds
	n := GetRedisSettings()
	conn, err := redis.DialURL(n)
	if err != nil {
		return ""
	}
	_, err = conn.Do("SETEX", sessionToken, "86400", name)
	if err != nil {
		// If there is an error in setting the cache, return an internal server error

		return ""
	}
	return sessionToken
}

func UpdateSession(w http.ResponseWriter, r *http.Request, sessionname string) {

	c, err := r.Cookie(sessionname)
	if err != nil {
		return
	}
	sessionToken := c.Value

	response, err := cache.Do("GET", sessionToken)
	if err != nil {

		return
	}
	if response == nil {

		return
	}
	// Now, create a new session token for the current user
	newsessionToken := SetSession(fmt.Sprintf("%s", response))

	// Delete the older session token
	_, err = cache.Do("DEL", sessionToken)
	if err != nil {

		return
	}

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    sessionname,
		Path:    "/",
		Value:   newsessionToken,
		Expires: time.Now().Add(86400 * time.Second),
	})

}

func DelSession(w http.ResponseWriter, r *http.Request, sessionname string) {
	fmt.Println(sessionname)
	c, err := r.Cookie(sessionname)
	if err != nil {

		return
	}
	sessionToken := c.Value

	http.SetCookie(w, &http.Cookie{
		Name:  sessionname,
		Path:  "/",
		Value: "",
	})

	response, err := cache.Do("GET", sessionToken)
	if err != nil {

		return
	}
	if response == nil {

		return
	}

	// Delete the older session token
	_, err = cache.Do("DEL", sessionToken)
	if err != nil {

		return
	}

}
