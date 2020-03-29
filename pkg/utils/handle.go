package utils

import (
	"log"
)

func HandleErr(err error, msg string) {
	if err != nil {
		log.Fatal(msg)
	}
}