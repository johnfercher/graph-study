package infra

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlConnection() (*sql.DB, error) {
	return sql.Open("mysql", "AdminUser:AdminPassword@tcp(127.0.0.1:3307)/RussianDoll")
}
