package services

var webhookURL string

func SetWebhook(url string) {
	webhookURL = url
}

func GetWebhook() string {
	return webhookURL
}
