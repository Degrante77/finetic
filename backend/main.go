
package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/Degrante77/finetic/config"
    "github.com/Degrante77/finetic/routes"
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
