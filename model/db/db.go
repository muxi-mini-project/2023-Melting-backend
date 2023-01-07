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
	if err != nil {
		log.Fatal(err)
	}
	var cert gomysql.Config
	tmp, _ := io.ReadAll(file)
	json.Unmarshal(tmp, &cert)
	DB, err = gorm.Open(mysql.Open(cert.FormatDSN()))
	if err != nil {
		log.Fatal(err)
	}

}
