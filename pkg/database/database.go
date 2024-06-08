package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Импорт с побочным эффектом для регистрации драйвера PostgreSQL
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database connection established")
}

func GetDB() *sql.DB {
	return db
}
