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
	"github.com/bradfitz/gomemcache/memcache"
	"strconv"
)

func UpdateRegDataFromtDB(Value string) {
	updatedData := [][]string{
		{"registration", Value},
	}
	UpdateRow("site_settings", updatedData, "id=1")
}

func ForbidRegistration() {
	UpdateRegDataFromtDB("0")
	n := GetMemcachedSettings()
	mc := memcache.New(n)
	mc.Set(&memcache.Item{Key: "registration", Value: []byte("0"), Expiration: 0})

}

func AllowRegistration() {
	UpdateRegDataFromtDB("1")
	n := GetMemcachedSettings()
	mc := memcache.New(n)
	mc.Set(&memcache.Item{Key: "registration", Value: []byte("1"), Expiration: 0})

}

func CheckRegistration() bool {

	n := GetMemcachedSettings()
	mc := memcache.New(n)
	it, err := mc.Get("registration")
	// Get multiple values
	//  it, err := mc.GetMulti([]string{"key_one", "key_two"})
	if err != nil {

		//There is no data in Mmecached
		//Should get from DB
		var n bool = GetSiteSettings("registration")
		return n
	}

	i, err := strconv.Atoi(string(it.Value))
	if err != nil {
		SetErrorLog(err.Error())
		return false
	}
	if i != 0 {
		return true
	} else {
		return false
	}

}
