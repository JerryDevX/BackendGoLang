package utils

import (
	"fmt"
	"os"
)


func ConfigType(t string) (string, error) {
	var url string
	switch t {
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
		case "mongo":
			// URL for MongoDB connection.
			url = fmt.Sprintf(
				"mongodb://%s",
				os.Getenv("MONGO_URL"),
			)	
		default:
			return "", fmt.Errorf("ConfigType: %s not found", t)
		}

	return url, nil
}
