package main

import (
	"challengeFinal/auth"
	"challengeFinal/database"
	"challengeFinal/models"
	"challengeFinal/records"
	"challengeFinal/shared"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

 func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(shared.Cors())
	records.AddRecordsRoutes(router)
	auth.AddAuthRoutes(router)
	return router
} 

func loadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .en file")
	}
}

func main() {
	loadEnvVars()
	database.CreateDbConnection()
	database.DBConn.AutoMigrate(&models.Record{}, &models.User{})
	router := SetupRouter()
	router.Run(":3333")
}
