//////////////////////////////////////////////////////////////////////
// mysql.go
//
// @usage
// 
//     1. Import this package.
//
//         --------------------------------------------------
//         import myMySQL "mysql"
//         --------------------------------------------------
//
//     2. Initialize database.
//
//         --------------------------------------------------
//         driverName := "DB_USER:DB_PASS@unix(/var/run/mysql/mysql.sock)/DB_NAME?parseTime=true"
//         myMySQL.Init(driverName)
//         --------------------------------------------------
//
//     3. Now, you can use it!!
//
//         3-1. When getting a connection.
//
//             --------------------------------------------------
//             db := myMySQL.Conn()
//             --------------------------------------------------
//
//
// MIT License
//
// Copyright (c) 2019 noknow.info
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A 
// PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTW//ARE.
//////////////////////////////////////////////////////////////////////
package mysql

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "log"
)

const (
    DRIVER_NAME = "mysql"
)

var (
    myDb *sql.DB
)


//////////////////////////////////////////////////////////////////////
// Initialize database.
//////////////////////////////////////////////////////////////////////
func Init(datasourceName string) {
    db, err := sql.Open(DRIVER_NAME, datasourceName)
    if err != nil {
        log.Fatalf("[FATAL] sql.Open() error: %s\n", err)
    }
    if err := db.Ping(); err != nil {
        db.Close()
        log.Fatalf("[FATAL] (*sql.DB) Ping() error: %s\n", err)
        return
    }
    myDb = db
}


//////////////////////////////////////////////////////////////////////
// Close database.
//////////////////////////////////////////////////////////////////////
func Close() {
    if myDb != nil {
        myDb.Close()
    }
}


//////////////////////////////////////////////////////////////////////
// Connect to the database.
//////////////////////////////////////////////////////////////////////
func Conn() *sql.DB {
    return myDb
}
