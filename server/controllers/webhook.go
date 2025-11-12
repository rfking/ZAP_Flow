package controllers

import (
    "github.com/labstack/echo/v4"
		"github.com/verbeux-ai/whatsmiau/repositories/instances"
		"github.com/verbeux-ai/whatsmiau/models"
)

var webhooks = make(map[string]string)

func AddWebhook(ctx echo.Context) error {
    instanceId := ctx.Param("instance")
    type Request struct {
        URL string `json:"url"`
    }
    var req Request
    if err := ctx.Bind(&req); err != nil {
        return ctx.JSON(400, echo.Map{"error": "request inválido"})
    }
    webhooks[instanceId] = req.URL
    return ctx.JSON(200, echo.Map{"message": "Webhook adicionado!", "url": req.URL})
}

func DeleteWebhook(ctx echo.Context) error {
    instanceId := ctx.Param("instance")
    delete(webhooks, instanceId)
    return ctx.JSON(200, echo.Map{"message": "Webhook removido!"})
}

func GetWebhook(ctx echo.Context) error {
	instanceID := ctx.Param("instance")

	instance, err := instances.FindByID(instanceID)
	if err != nil {
		return ctx.JSON(500, echo.Map{"error": "instância não encontrada"})
	}

	if instance.Webhook == nil {
		return ctx.JSON(200, echo.Map{"webhook": nil})
	}

	return ctx.JSON(200, echo.Map{"webhook": instance.Webhook})
}
