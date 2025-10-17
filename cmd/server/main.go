package main

import (
	"dept-collector/internal/api"
	"dept-collector/internal/config"

	_ "github.com/gin-gonic/gin"
)

func main() {

	db := config.ConnectDB()
	config.AutoMigrate(db)

	router := api.NewRouter(db)
	router.Run()
}
