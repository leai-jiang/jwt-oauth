package core

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

const (
	DriveName = "mysql"
	IP = "139.199.230.235"
	Port = "3306"
	User = "root"
	Password = "Jerusalem3"
	Database = "sparta"
)

var DB *sql.DB

func ConnectDB() {
	dataSourceName := strings.Join([]string{User, ":", Password, "@tcp(", IP, ":", Port, ")/", Database, "?charset=utf8&parseTime=True&loc=Local"}, "")
	DB, _ = sql.Open(DriveName, dataSourceName)

	DB.SetConnMaxLifetime(100)

	DB.SetMaxIdleConns(10)

	if err := DB.Ping();err != nil {
		fmt.Println(err)
		log.Panic(err)
	}
	log.Println("database connect success")
}

