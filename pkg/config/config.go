package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var  (
	
	// rest api server
	SERVICE_PORT string	= getenv("SERVICE_PORT", "2023")

	// job worker
	WORKER_QUOTA int
)

func init() {

	godotenv.Load()

	WORKER_QUOTA, _ = strconv.Atoi(getenv("WORKER_QUOTA", "3"))
}

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}