package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/asishshaji/thalir-backend/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateTables(db *gorm.DB) {

	prod := models.Product{}
	order := models.Order{}
	orderDetail := models.OrderDetail{}

	db.AutoMigrate(&prod, &order, &orderDetail)

}

func ConnectToDB(host, user, password, dbname string, port int) *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  psqlInfo,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	return db

}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return
	}
	log.Println("Loaded .env file ")

}
