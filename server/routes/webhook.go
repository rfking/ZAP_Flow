package main

import (
    "net/http"
    "github.com/verbeux-ai/whatsmiau/server/routes"
)

func main() {
    http.HandleFunc("/webhook", routes.WebhookHandler)
    http.ListenAndServe(":8081", nil)
}
