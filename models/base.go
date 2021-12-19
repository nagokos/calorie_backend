package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nagokos/calorie_backend/config"
)

const (
	tableNameUsers = "users"
)

var DbConnection *sql.DB

func init() {
	var err error
	DbConnection, err = sql.Open("mysql", config.Config.DBConf)
	if err != nil {
		log.Fatal("error:", err)
		return
	}

	cmd := fmt.Sprintf(`
	  CREATE TABLE IF NOT EXISTS %s (
			id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
			name VARCHAR(30) NOT NULL,
			email VARCHAR(256) UNIQUE NOT NULL,
		  password_hash VARCHAR(60),
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		)
	`, tableNameUsers)
	DbConnection.Exec(cmd)
}
