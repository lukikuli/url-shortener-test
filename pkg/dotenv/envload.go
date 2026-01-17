package dotenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Environment() {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Printf("error loading .env file: %s", err.Error())
		}
	} else {
		log.Println("running service without configuration from .env")
	}
}
