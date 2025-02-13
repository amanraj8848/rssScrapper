package auth

import (
	"errors"
	"net/http"
	"strings"
)

// getApiKey extracts the api key
// from the headers of an HTTP request
// Example: Authorization: ApiKey {actual api key}
func GetApiKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no api key found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("no api key found")
	}
	return vals[1], nil
}
