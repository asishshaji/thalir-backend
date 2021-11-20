package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/asishshaji/thalir-backend/controller"
	"github.com/asishshaji/thalir-backend/utils"

	repository "github.com/asishshaji/thalir-backend/repositories"
	"github.com/asishshaji/thalir-backend/services"
)

func main() {
	utils.LoadEnv()
	serverPort := os.Getenv("SERVER_PORT")

	dbHost := os.Getenv("HOST")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	envMode := os.Getenv("ENV_MODE")

	db := utils.ConnectToDB(dbHost, dbUsername, dbPassword, dbName, dbPort)

	utils.CreateTables(db, envMode)

	pR := repository.NewProductRepo(db)
	pS := services.NewProductService(pR)
	pC := controller.NewProductController(pS)

	oR := repository.NewOrderRepository(db)
	oS := services.NewOrderService(oR, pR)
	oC := controller.NewOrderController(oS)

	a := NewApp(fmt.Sprintf(":%s", serverPort), pC, oC)
	a.RunServer()
}
