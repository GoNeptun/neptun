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

package webserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	//"time"
	s "github.com/goneptune/neptune/semafor"
	sf "github.com/goneptune/neptune/systemfunctions"
	"github.com/gorilla/mux"
)

type MT struct {
	Live int `db:"live"`
}

type Lino struct {
	Tom string
}

func ExitApp(w http.ResponseWriter, r *http.Request) {
	var m string = "Neptune Stop \t"
	sf.SetLog(m)
	fmt.Printf(m)
	os.Exit(3)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	n := sf.LoadContentDirectory()
	favicon := n + "static/img/favicon.ico"
	if sf.FileExists(favicon) {
		http.ServeFile(w, r, favicon)
	} else {
		http.Redirect(w, r, "/404/", http.StatusFound)
	}
}

func RobotsTXTHandler(w http.ResponseWriter, r *http.Request) {
	n := sf.LoadContentDirectory()
	robots := n + "layouts/robots.txt"
	if sf.FileExists(robots) {
		http.ServeFile(w, r, robots)
	} else {
		w.Write([]byte("User-agent: * \nDisallow: / "))
	}
}

func SitemapHandler(w http.ResponseWriter, r *http.Request) {
	n := sf.LoadContentDirectory()
	sitemap := n + "static/sitemap.xml"
	if sf.FileExists(sitemap) {
		http.ServeFile(w, r, sitemap)
	} else {
		http.Redirect(w, r, "/404/", http.StatusFound)
	}
}

func StartService(port string) {
	//Initiate redis cache
	sf.InitCache()

	var m string = "Server Neptune Starting. Listen []:" + port + "\t"
	sf.SetLog(m)
	fmt.Printf(m)

	//Site pages
	//Open Pages
	n := sf.LoadContentDirectory()
	dir := n + "static/"
	cssdir := dir + "css/"
	imgdir := dir + "img/"
	r := mux.NewRouter()
	r.HandleFunc("/", s.IndexPage)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static",
		http.FileServer(http.Dir(dir))))
	r.PathPrefix("/css").Handler(http.StripPrefix("/css",
		http.FileServer(http.Dir(cssdir))))
	r.PathPrefix("/img").Handler(http.StripPrefix("/img",
		http.FileServer(http.Dir(imgdir))))

	r.NotFoundHandler = http.HandlerFunc(s.NotFoundAny)
	http.Handle("/", r)

	// http.HandleFunc("/openapi", s.Semafor)
	//http.HandleFunc("/install", s.Semafor)

	//  http.HandleFunc("/trade", s.Semafor)

	//API
	http.HandleFunc("/adminfuncs/", s.AdminFuncs)
	http.HandleFunc("/sysapi/", s.SysAPI)
	http.HandleFunc("/user/", s.UserAPI)

	//System
	http.HandleFunc("/exit", ExitApp)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/robots.txt", RobotsTXTHandler)
	http.HandleFunc("/sitemap.xml", SitemapHandler)
	http.HandleFunc("/404/", s.NotFound)
	http.HandleFunc("/logout/", s.LogOut)

	//Access Pages
	http.HandleFunc("/signup/", s.Signup)
	http.HandleFunc("/signin/", s.Signin)
	http.HandleFunc("/forgot/", s.ForgotPassword)

	//Admin pages
	http.HandleFunc("/admin/", s.Admin)

	//Pages for registered users
	http.HandleFunc("/profile/", s.Profile)
	http.HandleFunc("/profile/change_email/", s.ChangeEmail)
	http.HandleFunc("/profile/change_password/", s.ChangePassword)

	//Open Pages

	//Server start
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
