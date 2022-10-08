package main

import (
	"assignment2/database"
	_ "assignment2/docs"
	"assignment2/routers"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

const PORT = ":8080"

// @title           hacktiv8-go Assignment 2 API
// @version         1.0
// @description     This is a sample API for hacktiv8-go Assignment 2

// @contact.name   Indra Wijaya
// @contact.email  indra.wijaya2303@gmail.com

// @host      localhost:8080
// @BasePath  /
func main() {
	database.StartDB()

	r := routers.StartServer()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
