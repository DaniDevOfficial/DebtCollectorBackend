package main

import (
	"dept-collector/internal/api"

	_ "github.com/gin-gonic/gin"
)

func main() {
	router := api.NewRouter()
	router.Run()
}
