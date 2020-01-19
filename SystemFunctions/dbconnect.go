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

package SystemFunctions

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

func GetDBString() []string {
	n := LoadConfig("database")
	//n[0] = type, n[1] = host, n[2] = port, n[3] = dbname, n[4] = usr, n[5] = pwd, [n6] = protocol
	var database string = n[4] + ":" + n[5] + "@" + n[6] + "(" + n[1] + ":" + n[2] + ")/" + n[3]
	s := []string{n[0], database}
	//s:= [] string {"mysql", "root:root@tcp(127.0.0.1:8889)/exchange"}
	return s

}

func DBCheck() bool {
	n := GetDBString()
	//n[0] = type, n[1] = database
	db, err := sqlx.Open(n[0], n[1])
	if err != nil {
		SetErrorLog(err.Error())
	}
	err = db.Ping()
	if err != nil {
		return false
	}
	return true
}

func CountRows(value string, table string, condition string) int {
	var count int

	n := GetDBString()
	//n[0] = type, n[1] = database
	conn, err := sqlx.Connect(n[0], n[1])
	if err != nil {
		SetErrorLog(err.Error())
	}

	finalStatement := "SELECT COUNT(" + value + ") FROM `" + table + "` WHERE " + condition
	//fmt.Printf(finalStatement)
	err = conn.QueryRow(finalStatement).Scan(&count)
	switch {
	case err != nil:
		SetErrorLog("Error in CountRow function")
		return 0
	default:
		return count
	}
}

func SelectFrom(value string, table string, condition string) *sqlx.Rows {
	n := GetDBString()
	//n[0] = type, n[1] = database
	conn, err := sqlx.Connect(n[0], n[1])
	if err != nil {
		SetErrorLog(err.Error())
	}

	finalStatement := "SELECT " + value + " from " + table + " WHERE " + condition

	rows, err := conn.Queryx(finalStatement)
	if err != nil {
		SetErrorLog(err.Error())
	}
	return rows
}

func CreateTable() {
	//Создание таблицы
	n := GetDBString()
	//n[0] = type, n[1] = database
	conn, err := sqlx.Connect(n[0], n[1])
	if err != nil {
		SetErrorLog(err.Error())
	}
	var schema string = "CREATE TABLE `usersi` (`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,`name` varchar(255) NOT NULL)"
	conn.MustExec(schema)

}

func InsertRow(table string, name []string, value []string) {
	//Good
	n := GetDBString()
	//n[0] = type, n[1] = database
	conn, err := sqlx.Connect(n[0], n[1])
	if err != nil {
		SetErrorLog(err.Error())
	}

	NameString := strings.Join(name, "`, `")
	ValueString := strings.Join(value, "', '")

	finalStatement := "INSERT INTO " + table + " (`" + NameString + "`) VALUES ('" + ValueString + "')"

	res, err := conn.Exec(finalStatement)
	if err != nil {
		SetErrorLog(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		SetErrorLog(err.Error())
	}

	fmt.Printf("Created user with id:%d ", id)

}

func UpdateRow(table string, updatedData [][]string, condition string) {
	// Good
	n := GetDBString()
	//n[0] = type, n[1] = database
	conn, err := sqlx.Connect(n[0], n[1])
	if err != nil {
		SetErrorLog(err.Error())
	}

	var strs []string
	for _, v1 := range updatedData {
		s := strings.Join(v1, "=")
		strs = append(strs, s)

	}

	updatedDataString := strings.Join(strs, ", ")

	finalStatement := "UPDATE " + table + " set " + updatedDataString + " where " + condition
	fmt.Printf(finalStatement)

	_, err = conn.Exec(finalStatement)
	if err != nil {
		SetErrorLog(err.Error())
	}

}

func DeleteRow(table string, condition string) {
	// Good
	n := GetDBString()
	//n[0] = type, n[1] = database
	finalStatement := "DELETE FROM " + table + " where " + condition

	conn, err := sqlx.Connect(n[0], n[1])
	if err != nil {
		SetErrorLog(err.Error())
	}

	_, err = conn.Exec(finalStatement)
	if err != nil {
		SetErrorLog(err.Error())
	}

	//fmt.Printf("Row Deleted")
}
