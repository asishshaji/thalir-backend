package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/asishshaji/thalir-backend/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func CreateTables(db *bun.DB) {

	prod := models.Product{}
	order := models.Order{}
	orderHistory := models.OrderHistory{}
	db.NewCreateTable().Model(&prod).Exec(context.Background())
	db.NewCreateTable().Model(&orderHistory).Exec(context.Background())
	db.NewCreateTable().Model(&order).Exec(context.Background())

}

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
