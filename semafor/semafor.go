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
	sf "../system_functions"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  string `db:"age"`
}

type RequestStruct struct {
	Method string
}


func InsertIntoDB() {
	//Insert Row
	name := []string{"name", "age"}
	value := []string{"\"John\"", "32"}

	sf.InsertRow("usersi", name, value)
}

func UpdateDB() {
	//Update Row
	updatedData := [][]string{
		{"name", "\"Andy\""},
		{"age", "40"},
	}

	sf.UpdateRow("usersi", updatedData, "id=8")
}

func SelectFromDB() {
	//Select from DB
	SqlAnswer := sf.SelectFrom("*", "usersi", "id=8")

	var user User
	for SqlAnswer.Next() {

		err := SqlAnswer.StructScan(&user)
		if err != nil {
			panic(err)
		}
	}
	//For string %s and for Integer is %d if you do not know set %v

	n := "User name is: " + user.Name
	fmt.Printf(n)

}

//JSON request
func Semafor(w http.ResponseWriter, r *http.Request) {

	//make byte array
	out := make([]byte, 1024)

	bodyLen, err := r.Body.Read(out)
	if err != io.EOF {
		fmt.Println(err.Error())
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}
	var k RequestStruct
	err = json.Unmarshal(out[:bodyLen], &k)
	if err != nil {
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}
	fmt.Println(k.Method)
	w.Write([]byte(`{"error":"success"}`))
}




/*
//Simple request
func Semafor (w http.ResponseWriter, r *http.Request) {
  fmt.Println(r)
  b := make([]byte, 1024)
 fmt.Println(string(b)) //Data

  fmt.Println(r.Method) // JSON
  fmt.Println(string(strings.Join(r.Header["Content-Type"],""))) //POST
  w.Write([]byte("Hello Friends, Wecome!"))
}
*/
