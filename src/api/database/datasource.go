package database

import (
	"biller-api/src/api/config"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var once sync.Once

// GetDB access to DB
func GetDB() {
	once.Do(func() {
		connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", config.Global.DBUsername, config.Global.DBPassword, config.Global.DBHost, config.Global.DBName)

		DB = sqlx.MustConnect(config.Global.DBDriverName, connectionString+"&parseTime=true")
		DB.SetMaxIdleConns(100)
		DB.SetMaxOpenConns(400)
		DB.SetConnMaxLifetime(15 * time.Minute)
	})
}
