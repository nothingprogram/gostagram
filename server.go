package main

import (
	"fmt"
	"gostagram/app"
	"log"
	"net/http"
	"os"

	_ "gostagram/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func sanityCheck() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined...")
	}
}

// @title Gostagram Swagger API
// @version 1.0
// @host 127.0.0.1:4000
// @BasePath /api/v1
func main() {
	sanityCheck()
	e := app.Routes()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	log.Printf("Listening %s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT"))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}
