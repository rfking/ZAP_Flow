package routes

import (
    "io/ioutil"
    "log"
    "net/http"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
    log.Printf("Recebido %s no webhook", r.Method)
    body, _ := ioutil.ReadAll(r.Body)
    log.Println("Headers:", r.Header)
    log.Println("Payload:", string(body))
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status": "success"}`))
}
