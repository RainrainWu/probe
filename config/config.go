package config

import (
	"os"
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

var  (
	
	// rest api server
	SERVICE_PORT string	= getenv("SERVICE_PORT", "2023")

	// job worker
	WORKER_QUOTA int = strToInt(getenv("WORKER_QOUTA", "3"))

	// jwt auth
	JWT_SECRET string = getenv("JWT_SECRET", "mysecret")
	USERNAME string = getenv("USERNAME", "probeuser")
	PASSWORD string = getenv("PASSWORD", "probepass")

	// zap logger
	LOG_FILE string = getenv("LOG_FILE", "./logs/probe.log")
)

func init() {

	godotenv.Load()
}

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func strToInt(source string) int {
	result, err := strconv.Atoi(source)
	if err != nil {
		log.Fatal("Error when parse " + source + " to int")
	}
	return result
}
