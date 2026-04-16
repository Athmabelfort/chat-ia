package controllers

import (
	"context"
	"os"

	"github.com/chat-ia/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/genai"
)

var ctx = context.Background()

func GetMessage(c *gin.Context) {
	var message models.MessageIA
	error := c.ShouldBindJSON(&message)
	if error != nil {
		c.JSON(400, gin.H{"error": "Requisição inválida para a rota"})
		return
	}

	key := os.Getenv("GEMINI_API_KEY")
	if key == "" {
		c.JSON(500, gin.H{"error": "GEMINI_API_KEY não definida"})
		return
	}

	config := &genai.ClientConfig{
		APIKey: key,
	}
	client, err := genai.NewClient(ctx, config)
	if err != nil {
		c.JSON(500, gin.H{"erro": "Erro ao criar o client do Gemini"})
		return
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text(message.Message), nil,
	)
	if err != nil {
		c.JSON(500, gin.H{"erro": "Erro ao gerar conteúdo"})
		return
	}

	if len(result.Candidates) == 0 ||
		len(result.Candidates[0].Content.Parts) == 0 {
		c.JSON(500, gin.H{"error": "Nenhum conteúdo foi gerado após o prompt"})
		return
	}
	text := result.Candidates[0].Content.Parts[0].Text
	c.JSON(200, gin.H{"message": string(text)})
}
