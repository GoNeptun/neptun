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
	viper "github.com/spf13/viper"
)

type Params map[string]interface{}
type Menus map[string]interface{}
type Data map[string]interface{}

func GetSS(lg string, param string) string {
	n := GetSiteSettingsPath()
	y := viper.New()
	y.SetConfigName(n[1]) // name of config file (without extension)
	y.AddConfigPath(n[0])
	err := y.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file

		SetErrorLog("Please Check config file. It is an error in it...")
	}

	var sfg string = y.GetString(param)
	return sfg
}

func ParamsConfig(lg string) Params {
	n := GetSiteSettingsPath()
	y := viper.New()
	y.SetConfigName(n[1]) // name of config file (without extension)
	y.AddConfigPath(n[0])
	err := y.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file

		SetErrorLog("Please Check config file. It is an error in it...")
	}

	var sfg Params
	check := y.IsSet("params")
	if check {
		prod := y.Sub("params")
		err = prod.Unmarshal(&sfg)
		if err != nil { // Handle errors reading the config file

			SetErrorLog(err.Error())
		}
	}
	//fmt.Printf(sfg["description"].(string))
	//Add Languages
	lang := "Languages." + lg
	check = y.IsSet(lang)
	if check {
		k := y.Sub(lang)
		err = k.Unmarshal(&sfg)
		if err != nil { // Handle errors reading the config file

			SetErrorLog(err.Error())
		}
	}
	return sfg

}

func MenusConfig(lg string) Menus {
	n := GetSiteSettingsPath()
	y := viper.New()
	y.SetConfigName(n[1]) // name of config file (without extension)
	y.AddConfigPath(n[0])
	err := y.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file

		SetErrorLog("Please Check config file. It is an error in it...")
	}

	var sfg Menus

	check := y.IsSet("menu")
	if check {
		prod := y.Sub("menu")
		err = prod.Unmarshal(&sfg)
		if err != nil { // Handle errors reading the config file

			SetErrorLog(err.Error())
		}
	}
	//fmt.Printf(sfg["description"].(string))
	//Add Languages

	lang := "Languages." + lg + ".menu"
	check = y.IsSet(lang)
	if check {
		k := y.Sub(lang)
		err = k.Unmarshal(&sfg)
		if err != nil { // Handle errors reading the config file

			SetErrorLog(err.Error())
		}
	}
	//fmt.Printf(sfg["postpend"].(string))

	return sfg

}
