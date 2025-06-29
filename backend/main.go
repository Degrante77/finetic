
package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "yourmodule/config"
    "yourmodule/routes"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    config.ConnectDatabase()

    router := gin.Default()
    routes.SetupRoutes(router)

    port := os.Getenv("PORT")
    if port == ""{
        port = "8080"
    }

    router.Run(":" + port)
}
