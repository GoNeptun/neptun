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
	viper "github.com/spf13/viper"
	"os"
	"strings"
)

func HashConfigPasswords() {
	viper.SetConfigName(configname) // name of config file (without extension)
	viper.AddConfigPath(configpath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file

		SetErrorLog("Please Check config file. It is an error in it...")
	}
	var pwd string = viper.Get("database.password").(string)
	//Hashed Password started with $2a##
	s := strings.Split(pwd, "a##")
	if s[0] != "$2" {
		// The password is not hashed

		// encrypt password
		encryptMsg, err := encrypt(pwd)
		if err != nil {
			SetErrorLog(err.Error())
		}
		//Create new passwrd recort
		var newpasswordrecord interface{} = "$2a##" + encryptMsg
		//s := make(interface{},  newpasswordrecord)
		viper.Set("database.password", newpasswordrecord)
		viper.WriteConfig()
	}

	var emailpwd string = viper.Get("email.password").(string)
	semail := strings.Split(emailpwd, "a##")
	if semail[0] != "$2" {
		// The password is not hashed

		// encrypt password
		encryptMsg, err := encrypt(emailpwd)
		if err != nil {
			SetErrorLog(err.Error())
		}
		//Create new passwrd recort
		var newpasswordrecord interface{} = "$2a##" + encryptMsg
		//s := make(interface{},  newpasswordrecord)
		viper.Set("email.password", newpasswordrecord)
		viper.WriteConfig()
	}

}

func DecryptConfigPasswords(pwd string) string {
	s := strings.Split(pwd, "a##")
	if s[0] != "$2" {
		return pwd
	} else {
		msg, _ := decrypt(s[1])
		return msg

	}
}

func CheckConfig() {
	viper.SetConfigName(configname) // name of config file (without extension)
	viper.AddConfigPath(configpath)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			var m string = "Config file was not found... \t"
			fmt.Printf(m)
			SetErrorLog(m)
			SetLog("Server stop")
			os.Exit(3)

		} else {
			// Config file was found but another error was produced
			var m string = "Please Check config file. It is an error in it... \t"
			fmt.Printf(m)
			SetErrorLog(m)
			SetLog("Server stop")
			os.Exit(3)
		}
	}
	//Хэширование паролей
	HashConfigPasswords()

	SetLog("Config file... OK")
}

func LoadConfig(conf string) []string {
	viper.SetConfigName(configname) // name of config file (without extension)
	viper.AddConfigPath(configpath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file

		SetErrorLog("Please Check config file. It is an error in it...")
	}
	switch os := conf; os {
	case "server":
		var port string = viper.Get("server.port").(string)
		var ip string = viper.Get("server.ip").(string)
		var domain string = viper.Get("server.baseurl").(string)
		if domain == "" {
			domain = "http://" + ip + ":" + port
		}
		var logsdir string = viper.Get("server.logs_dir").(string)
		var contentdir string = viper.Get("server.content_dir").(string)
		//var theme string = viper.Get("server.theme").(string)
		//var languageCode string = viper.Get("server.languageCode").(string)
		//var title string = viper.Get("server.title").(string)
		//var googleAnalytics string = viper.Get("server.googleAnalytics").(string)
		s := []string{port, logsdir, domain, ip, contentdir}
		return s

	case "database":
		var typee string = viper.Get("database.type").(string)
		var host string = viper.Get("database.host").(string)
		var port string = viper.Get("database.port").(string)
		var dbname string = viper.Get("database.database_name").(string)
		var usr string = viper.Get("database.user").(string)
		var pwd string = viper.Get("database.password").(string)
		var epwd string = DecryptConfigPasswords(pwd)
		var protocol string = viper.Get("database.protocol").(string)
		s := []string{typee, host, port, dbname, usr, epwd, protocol}
		return s

	case "memcached":
		var host string = viper.Get("memcached.host").(string)
		var port string = viper.Get("memcached.port").(string)

		s := []string{host, port}
		return s

	case "redis":
		var host string = viper.Get("redis.host").(string)
		s := []string{host}
		return s

	case "email":
		var host string = viper.Get("email.host").(string)
		var port string = viper.Get("email.port").(string)
		var euser string = viper.Get("email.user").(string)
		var mailpwd string = viper.Get("email.password").(string)
		var emailpwd string = DecryptConfigPasswords(mailpwd)
		var eaddress string = viper.Get("email.address").(string)
		var replyto string = viper.Get("email.reply").(string)
		s := []string{host, port, euser, emailpwd, eaddress, replyto}
		return s

	default:

		return nil
	}

}

func GetCustomCSS() []string {
	n := GetSiteSettingsPath()
	y := viper.New()
	y.SetConfigName(n[1]) // name of config file (without extension)
	y.AddConfigPath(n[0])
	err := y.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file

		SetErrorLog("Please Check config file. It is an error in it...")
	}
	var customcss []string = y.GetStringSlice("custom_css")
	return customcss
}

func GetCustomJS() []string {
	n := GetSiteSettingsPath()
	y := viper.New()
	y.SetConfigName(n[1]) // name of config file (without extension)
	y.AddConfigPath(n[0])
	err := y.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file

		SetErrorLog("Please Check config file. It is an error in it...")
	}
	var customjs []string = y.GetStringSlice("custom_js")
	return customjs
}

func GetDisableFonts() bool {
	n := GetSiteSettingsPath()
	y := viper.New()
	y.SetConfigName(n[1]) // name of config file (without extension)
	y.AddConfigPath(n[0])
	err := y.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file

		SetErrorLog("Please Check config file. It is an error in it...")
	}
	var disableFonts bool = y.GetBool("disableFonts")
	return disableFonts
}

func GetSiteSettingsPath() []string {
	viper.SetConfigName(configname) // name of config file (without extension)
	viper.AddConfigPath(configpath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file

		SetErrorLog("Please Check config file. It is an error in it...")
	}
	var dir string = viper.GetString("server.site_config_dir")
	var name string = viper.GetString("server.site_config_name")
	s := []string{dir, name}
	return s
}
