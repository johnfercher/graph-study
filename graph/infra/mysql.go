package infra

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlConnection() (*sql.DB, error) {
	//return sql.Open("mysql", "AdminUser:AdminPassword@tcp(mysql:3306)/Graph")
	return sql.Open("mysql", "AdminUser:AdminPassword@tcp(localhost:3306)/Graph")
}
