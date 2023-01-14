package db

import (
	"encoding/json"
	"io"
	"log"
	"os"

	gomysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDB() {
	var err error
	file, err := os.Open("./conf/db.json")
	var cert gomysql.Config
	if err != nil {
		cert = gomysql.Config{
			User:                 os.Getenv("user"),
			Passwd:               os.Getenv("pwd"),
			Net:                  "tcp",
			Addr:                 os.Getenv("addr"),
			DBName:               os.Getenv("db"),
			AllowNativePasswords: true,
		}
	} else {
		tmp, _ := io.ReadAll(file)
		json.Unmarshal(tmp, &cert)
	}
	DB, err = gorm.Open(mysql.Open(cert.FormatDSN()))
	if err != nil {
		log.Fatal(err)
	}

}
