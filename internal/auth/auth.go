package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApi from header function
// Gets header value of key --> Authorization
// Split the values --> expects len = 2
// Returns error if wrong format else
// Return apikey value
func GetApiKey(header http.Header) (string, error){
	value := header.Get("Authorization")
	if value == "" {
		return "", errors.New("no authentication info found")
	}
	
	values := strings.Split(value, " ")
	if len(values) != 2{
		return "", errors.New("malformed authentication header")
	}
	if values[0] != "ApiKey"{
		return"", errors.New("malformed auth header api-key")
	}

	return values[1], nil
}