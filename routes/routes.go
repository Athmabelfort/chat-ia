package routes

import (
	"fmt"

	"github.com/chat-ia/controllers"
	"github.com/gin-gonic/gin"
)

func SettingupRoutes() {
	r := gin.Default()

	r.POST("/message", controllers.GetMessage)

	error := r.Run(":8080")

	if error != nil {
		fmt.Println("Erro ao iniciar o servidor na porta 8080: ", error)
	}
}
