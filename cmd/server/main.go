package main

import "github.com/gin-gonic/gin"

func main() {
	router := api.NewRouter()
	router.Run()
}
