package helpers

import (
	"net/http"
	"strings"
)

func GetApiKey(request *http.Request) (string, error) {
	authHeader := request.Header.Get("Authorization")
	apiKey := strings.Replace(authHeader, "ApiKey ", "", 1)

	return apiKey, nil
}
