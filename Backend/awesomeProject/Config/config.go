package Config

import (
	"database/sql"
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "12345678"
	dbName := "ClinicApp"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(mysql-container:3306)/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	return db
}
