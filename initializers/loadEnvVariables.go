package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// var port, userDB, passDB, host string

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// port = os.Getenv("PORT")
	// userDB = os.Getenv("USER_DB")
	// passDB = os.Getenv("PASSWORD_DB")
	// host = os.Getenv("HOST")
}