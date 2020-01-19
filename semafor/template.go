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
	"bytes"
	"fmt"
	"github.com/gomarkdown/markdown"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func partial(str string, data *HTMLData) template.HTML {

	funcMap := template.FuncMap{
		"ifIE": func() template.HTML {
			return template.HTML("<!--[if IE]>")
		},
		"endif": func() template.HTML {
			return template.HTML("<![endif]-->")
		},
		"safeHTML":    safeHTML,
		"partial":     partial,
		"markdownify": Markdownify,
		"relURL":      relURL,
		"absURL":      relURL,
	}

	//Read the file
	n := sf.LoadContentDirectory()

	dir := n + "layouts/partials/" + str

	b, err := ioutil.ReadFile(dir) // just pass the file name
	if err != nil {
		sf.SetErrorLog(err.Error())
	}

	t, err := template.New("main").Funcs(funcMap).Parse(string(b))
	if err != nil {
		sf.SetErrorLog(err.Error())
		return template.HTML("")
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "main", data); err != nil {
		sf.SetErrorLog(err.Error())
		return template.HTML("")
	}
	result := tpl.String()

	return template.HTML(result)
}

func relURL(str string) template.HTML {
	n := sf.BaseURL()
	ans := n + str
	return template.HTML(ans)
}

func safeHTML(str string) template.HTML {
	return template.HTML(str)
}

func Markdownify(data interface{}) template.HTML {
	md := []byte(fmt.Sprintf("%v", data))
	output := markdown.ToHTML(md, nil, nil)

	m := template.HTML(output)
	return m

}

func (data *HTMLData) ContentToData(r *http.Request) {

	funcMap := template.FuncMap{
		"ifIE": func() template.HTML {
			return template.HTML("<!--[if IE]>")
		},
		"endif": func() template.HTML {
			return template.HTML("<![endif]-->")
		},
		"safeHTML":    safeHTML,
		"partial":     partial,
		"markdownify": Markdownify,
		"relURL":      relURL,
	}

	//Read the file
	n := sf.LoadContentDirectory()
	urlpath := strings.Split(r.URL.Path, "/")
	ln := len(urlpath)
	ind := ln - 2
	fl := urlpath[ind]
	var fls string
	if fl != "" {
		fls = fl
	} else {
		fls = "main"
	}

	dir := n + "layouts/content/" + fls + ".html"

	if fileExists(dir) {
		b, err := ioutil.ReadFile(dir) // just pass the file name
		if err != nil {
			sf.SetErrorLog(err.Error())
		}

		t, err := template.New("main").Funcs(funcMap).Parse(string(b))
		if err != nil {
			sf.SetErrorLog(err.Error())

		}

		var tpl bytes.Buffer
		if err := t.ExecuteTemplate(&tpl, "main", data); err != nil {
			sf.SetErrorLog(err.Error())

		}
		result := tpl.String()

		ans := template.HTML(result)
		data.Site.Content = ans
	} else {
		data.Site.Content = ""
	}
}

func (data *HTMLData) ShowPage(w http.ResponseWriter, r *http.Request, page string) {
	data.StandartInfo()
	data.ContentToData(r)
	n := sf.LoadContentDirectory()

	index_dir := n + "layouts/"

	pg := ""
	switch page {
	case "service":
		pg = "service.html"
	case "maintenance":
		pg = "maintenance.html"
	default:
		pg = "index.html"
	}

	var index_page = path.Join(index_dir, pg)

	//Check for registrtion
	if !data.Menu.Login {
		var rs = sf.CheckRegistration()
		data.RegDataToHTML(rs)
	}

	funcMap := template.FuncMap{
		"ifIE": func() template.HTML {
			return template.HTML("<!--[if IE]>")
		},
		"endif": func() template.HTML {
			return template.HTML("<![endif]-->")
		},
		"safeHTML":    safeHTML,
		"partial":     partial,
		"markdownify": Markdownify,
		"relURL":      relURL,
	}

	zzz, err := ioutil.ReadFile(index_page) // just pass the file name
	if err != nil {
		sf.SetErrorLog(err.Error())
	}
	t, err := template.New("main").Funcs(funcMap).Parse(string(zzz))

	if err != nil {
		sf.SetErrorLog(err.Error())
		http.Error(w, "500 Internal Server Error", 500)
		return
	}

	if err := t.ExecuteTemplate(w, "main", data); err != nil {
		sf.SetErrorLog(err.Error())
		http.Error(w, "500 Internal Server Error", 500)
		return
	}

}
