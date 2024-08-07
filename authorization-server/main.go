package main

import (
	"fmt"
	"log"

	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/client_app"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/config"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/database"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/user"
)

func main() {
	log.Println("Loading configuration...")
	err := config.Init()
	if err != nil {
		log.Fatal("Error loading configuration", err)
	}
	log.Println("Configuration loaded")
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}
	err = db.AutoMigrate(&user.User{}, &client_app.ClientApp{})
	if err != nil {
		log.Fatal("Auto migration failed", err)
	}
	log.Println("Database connected and migrated successfully")
	server := createServer()
	server.Run(fmt.Sprintf(":%s", config.PORT))
	log.Println("Running authorization-server at", fmt.Sprintf("0.0.0.0:%s", config.PORT))
}
