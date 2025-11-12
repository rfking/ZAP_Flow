package routes

import (
    "github.com/labstack/echo/v4"
    "github.com/verbeux-ai/whatsmiau/server/controllers"
)

func Webhook(group *echo.Group) {
    group.POST("", controllers.AddWebhook)
    group.DELETE("", controllers.DeleteWebhook)
}
