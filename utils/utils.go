package utils

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func ConnectToDB(host, user, password, dbname string, port int) *bun.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	pgdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(pgdb, pgdialect.New())

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return db

}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return
	}
	log.Println("Loaded .env file ")

}
