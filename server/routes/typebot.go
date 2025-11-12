package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Handler para configurar o typebot na API do WhatsApp
func SetTypebotHandler(c *gin.Context) {
    instance := c.Param("instance")

    // Lógica para configurar o typebot usando o valor de instance
    // Aqui você pode adicionar sua integração real com o Typebot e o WhatsApp
    // Exemplo de resposta de sucesso:
    c.JSON(http.StatusOK, gin.H{
        "message": "Typebot configurado com sucesso!",
        "instance": instance,
    })
}

// Função para registrar a rota
func RegisterTypebotRoutes(router *gin.Engine) {
    router.POST("/typebot/set/:instance", SetTypebotHandler)
}
