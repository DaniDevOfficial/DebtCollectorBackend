package main

import (
	"dept-collector/internal/api"
	"dept-collector/internal/config"
	"dept-collector/internal/pkg/validator"

	_ "github.com/gin-gonic/gin"
)

// @title           DeptCollector
// @version         Alpha
// @description     API for the DebtCollector app
// @host            localhost:8080
// @BasePath        /api
func main() {

	db := config.ConnectDB()
	//config.AutoMigrate(db)

	validator.InitCustomValidators()

	router := api.NewRouter(db)
	router.Run()
}
