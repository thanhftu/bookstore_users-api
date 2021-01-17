package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	mysqlUsersUsername = "mysql_users_username"
// 	mysqlUsersPassword = "mysql_users_password"
// 	mysqlUsersHost     = "mysql_users_host"
// 	mysqlUsersSchema   = "mysql_users_schema"
// )

var (
	// Client connect to data
	Client *sql.DB
	// username = os.Getenv(mysqlUsersUsername)
	// password = os.Getenv(mysqlUsersPassword)
	// host     = os.Getenv(mysqlUsersHost)
	// schema   = os.Getenv(mysqlUsersSchema)
)

func init() {

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "hoangsam", "127.0.0.1:3306", "users_db",
	)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
