package db

import (
	"database/sql"

	"github.com/basic-echo-golang/config"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

// Inisiasi database
func Init() {
	conf := config.GetConfig()

	// format username:password@protocol(address:port)/dbname?param=value
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	// Membuka koneksi ke mysql
	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("connectionString error")
	}

	// cek apakah connectionString sudah ok atau belum
	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

// Return DB
func CreateCon() *sql.DB {
	return db
}
