package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/verbeux-ai/whatsmiau/server/services"
)

type WebhookReq struct {
	Url string `json:"url"`
}

func SetWebhook(ctx echo.Context) error {
	var body WebhookReq

	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
	}

	services.SetWebhook(body.Url)

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "webhook set",
		"url":     body.Url,
	})
}

func GetWebhook(ctx echo.Context) error {
	url := services.GetWebhook()

	return ctx.JSON(http.StatusOK, echo.Map{
		"url": url,
	})
}

func DeleteWebhook(ctx echo.Context) error {
	services.DeleteWebhook()
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "webhook deleted",
	})
}
