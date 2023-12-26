package helper

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDbConn() *sqlx.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PSWD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var once sync.Once
	var dbConn *sqlx.DB

	once.Do(func() {
		db, err := sqlx.Connect("postgres", psqlInfo)
		if err != nil {
			log.Fatalln(err)
		}
		dbConn = db

		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(1)

		if err := db.Ping(); err != nil {
			log.Fatalln(err)
		}

	})

	if dbConn != nil {
		log.Println("database connected")
	} else {
		log.Println("database not connected")
	}

	return dbConn
}
