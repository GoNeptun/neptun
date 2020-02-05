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

package main

import (
	WebServer "./webserver"
	"fmt"
	"os"
	//"time"
	sf "./systemfunctions"
)

func main() {

	sf.CheckConfig()

	sf.SetLog("Neptune Starting")
	sf.SetLog("Check config file...")

	sf.SetLog("Check Database connection...")

	if sf.DBCheck() {
		sf.SetLog("Database connection... OK")

	} else {
		sf.SetLog("DataBase Connection Error")
		sf.SetLog("Server Stop")
		fmt.Printf("DataBase Connection Error \t")
		fmt.Printf("Server Stop \t")
		os.Exit(3)
	}

	n := sf.LoadConfig("server")

	if n[0] != "" {
		WebServer.StartService(n[0])
	} else {
		WebServer.StartService("8090")
	}

}
