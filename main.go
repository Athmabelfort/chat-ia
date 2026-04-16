package main

import (
	"fmt"

	"github.com/chat-ia/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Aviso: não foi possível carregar .env:", err)
	}

	routes.SettingupRoutes()
}
