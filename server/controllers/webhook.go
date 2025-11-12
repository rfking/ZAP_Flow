package controllers

import (
    "github.com/labstack/echo/v4"
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
