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

func LoadContentDirectory() string {
	n := LoadConfig("server")
	m := GetSS("en", "theme")
	r := n[4] + "themes/" + m + "/"
	return r

}

func LoadTitle() string {
	n := ParamsConfig("en")
	return n["title"].(string)
}

func BaseURL() string {
	n := LoadConfig("server")
	baseurl := n[2] + "/"
	return baseurl
}

func LangCode() string {
	n := GetSS("en", "languageCode")
	return n
}
