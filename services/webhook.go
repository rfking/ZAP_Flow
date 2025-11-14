package services

import "sync"

var (
	webhookURL string
	lock       sync.RWMutex
)

func SetWebhook(url string) {
	lock.Lock()
	webhookURL = url
	lock.Unlock()
}

func GetWebhook() string {
	lock.RLock()
	defer lock.RUnlock()
	return webhookURL
}

func DeleteWebhook() {
	lock.Lock()
	webhookURL = ""
	lock.Unlock()
}
