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
	"github.com/bradfitz/gomemcache/memcache"
	"strconv"
)

type MailConfDB struct {
	Mail int `db:"mail_confirmed"`
}

func UpdateMailDataFromtDB(Value string) {
	updatedData := [][]string{
		{"mail_confirmation", Value},
	}
	UpdateRow("site_settings", updatedData, "id=1")
}

func OffMailConfirmation() {
	UpdateMailDataFromtDB("0")
	n := GetMemcachedSettings()
	mc := memcache.New(n)
	mc.Set(&memcache.Item{Key: "mail_confirmation", Value: []byte("0"), Expiration: 0})

}

func OnMailConfirmation() {
	UpdateMailDataFromtDB("1")
	n := GetMemcachedSettings()
	mc := memcache.New(n)
	mc.Set(&memcache.Item{Key: "mail_confirmation", Value: []byte("1"), Expiration: 0})

}

func CheckMailConfStatus() bool {

	n := GetMemcachedSettings()
	mc := memcache.New(n)
	it, err := mc.Get("mail_confirmation")
	// Get multiple values
	//  it, err := mc.GetMulti([]string{"key_one", "key_two"})
	if err != nil {

		//There is no data in Mmecached
		//Should get from DB
		var n bool = GetSiteSettings("mail_confirmation")
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

func UserMailConf(uname string) bool {
	var m bool = CheckMailConfStatus()
	if m {
		//Check does user confirm his email
		n := GetMemcachedSettings()
		mc := memcache.New(n)
		it, err := mc.Get(uname)
		if err != nil {

			//There is no data in Mmecached
			//Should get from DB
			n := "`name` = '" + uname + "'"
			SqlAnswer := SelectFrom("mail_confirmed", "users", n)
			var user MailConfDB
			for SqlAnswer.Next() {

				err := SqlAnswer.StructScan(&user)
				if err != nil {

					SetErrorLog(err.Error())
				}
			}

			if user.Mail != 0 {
				mc.Set(&memcache.Item{Key: uname, Value: []byte("1"), Expiration: 0})

				return true
			} else {
				return false
			}
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

	} else {
		return true
	}
}
