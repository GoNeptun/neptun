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
	sf "../SystemFunctions"
	"html/template"
)

type error interface {
	Error() string
}

type SysData struct {
	Title   string
	Content string
	Login   bool
	Admin   bool
}

type HTMLData struct {
	Header       HeaderData
	SysMsg       SysMsgData
	Menu         MenuData
	Page         template.HTML
	Content      ContentData
	HTMLContent  HTMLContentData
	Site         SiteStruct
	IsTranslated bool
}

type SiteStruct struct {
	LanguageCode string
	Title        string
	BaseURL      template.URL
	Params       sf.Params
	Language     Lang
	Menus        sf.Menus
	Data         sf.Data
	Content      template.HTML
}

type Lang struct {
	Lang string
}

//type Params map[string]interface{}
/*
type ParamsStruct struct {
	description string
	name string
	custom_css []string
	custom_js []string
	disableFonts bool

}
*/
type HeaderData struct {
	Done bool
	JS   Jsdata
	// CSS, JS string
}

type Jsdata struct {
	Done bool
	Body []string
}

type SysMsgData struct {
	Done    bool
	Message string
}

type MenuData struct {
	Done           bool
	Login, Admin   bool
	IsRegistration bool
}

type ContentData struct {
	Done  bool
	Title string
	Body  []string
}

type HTMLContentData struct {
	Done  bool
	Body  []template.HTML
	Title template.HTML
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type ForgotPass struct {
	Mail string `json:"mail", db:"mail"`
}

type UsersTable struct {
	UID            string `db:"id"`
	Name           string `db:"name"`
	Password       string `db:"pass"`
	Mail           string `db:"mail"`
	Mail_Confirmed int    `db:"mail_confirmed"`
	MailSent       string `db:"mailsent"`
	Created        string `db:"created"`
	Access         string `db:"access"`
	Login          string `db:"login"`
	Status         int    `db:"status"`
	Completed      int    `db:"completed"`
}

type TimeHash struct {
	Mail     string `db:"mail"`
	Name     string `db:"name"`
	Deadline int    `db:"deadline"`
}

type SiteSettingsTable struct {
	ID               int `db:"id"`
	Maintenance      int `db:"maintenance"`
	MailConfirmation int `db:"mail_confirmation"`
	Registration     int `db:"registration"`
}

type Roles struct {
	Admin string `db:"admin"`
}

type SignUpCred struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"name"`
	Mail     string `json:"mail", db:"mail"`
}
