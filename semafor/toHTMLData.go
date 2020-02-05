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
	sf "github.com/goneptune/neptune/systemfunctions"
	"html/template"
	//"fmt"
)

func (r *HTMLData) MenuToHTML(Login bool, Admin bool) {

	r.Menu = MenuData{
		Done:  true,
		Login: Login,
		Admin: Admin,
	}

}

func (r *HTMLData) RegDataToHTML(reg bool) {

	r.Menu = MenuData{
		Done:           true,
		IsRegistration: reg,
	}

}

func (r *HTMLData) HeaderToHTML(Title string) {

	n := sf.LoadTitle()
	fullTitle := Title + " | " + n
	r.Site = SiteStruct{
		Title: fullTitle,
	}

}

func (r *HTMLData) BodyToHTML(body []string) {

	r.Content = ContentData{
		Done:  true,
		Body:  append(body),
		Title: body[0],
	}

}

func (r *HTMLData) HTMLBodyToHTML(body []template.HTML) {

	r.HTMLContent = HTMLContentData{
		Done:  true,
		Body:  append(body),
		Title: body[0],
	}

}

func (r *HTMLData) JstoHtml(files []string) {

	r.Header.JS = Jsdata{
		Done: true,
		Body: append(files),
	}
}

func (r *HTMLData) StandartInfo() {

	r.Site = SiteStruct{
		LanguageCode: sf.LangCode(),
		BaseURL:      template.URL(sf.BaseURL()),
		Title:        r.Site.Title,
	}

	r.Site.Language = Lang{
		Lang: "en",
	}

	r.Site.Params = sf.ParamsConfig(string(r.Site.Language.Lang))
	r.IsTranslated = true
	r.Site.Menus = sf.MenusConfig(string(r.Site.Language.Lang))
}
