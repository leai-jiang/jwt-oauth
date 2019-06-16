package core

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sparta/config"
	"strings"
)

var DB *sql.DB

func ConnectDB() {
	var (
		DriveName = "mysql"
		IP = config.Viper.GetString("mysql.host")
		Port = config.Viper.GetString("mysql.port")
		User = config.Viper.GetString("mysql.user")
		Password = config.Viper.GetString("mysql.password")
		Database = config.Viper.GetString("mysql.database")
	)
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

