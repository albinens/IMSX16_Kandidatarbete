package utils

import "os"

func IsProduction() bool {
	return os.Getenv("ENVIRONMENT") == "production"
}
