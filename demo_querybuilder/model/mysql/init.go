package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	MysqlDriver   = "mysql"
	MysqlUser     = "namnt"
	MysqlPassWord = "123456"
	MysqlDbName   = "bg_student"
	MysqlHost     = "localhost" //imp.bidgear.com
	MysqlPort     = 3306
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", MysqlUser, MysqlPassWord, MysqlHost, MysqlPort, MysqlDbName)
var MySQL *sql.DB

func Connect() error {
	fmt.Println("dsn...", dsn)
	var err error
	MySQL, err = sql.Open(MysqlDriver, dsn)

	return err
}

func ContinueConnectMySQL() {
	err := MySQL.Ping()
	if err != nil {
		fmt.Println(err.Error())
		MySQL, _ = sql.Open(MysqlDriver, dsn)
	}
}

func Close() error {
	err := MySQL.Close()
	return err
}
