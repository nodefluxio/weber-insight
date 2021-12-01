package main

import (
	"os"
	"weber-insight/controllers"
	"weber-insight/database"
	"weber-insight/models"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	// Initialize DB Connection
	db := database.InitDB()

	// Create Model
	model := models.New(db)

	// Create Controller
	ctrl := controllers.New(model)

	r := setupRouter(ctrl)
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")

	_ = r.Run(host + ":" + port)
}
