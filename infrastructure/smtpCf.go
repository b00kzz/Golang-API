package infrastructure

import (
	"os"

	"github.com/joho/godotenv"
)

func EnvSMTP_ADDR() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("SMTP_ADDR")
}
func EnvSMTP_From() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("SMTP_FROM")
}
func EnvSMTP_User() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("SMTP_USER")
}
func EnvSMTP_Password() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("SMTP_PASSWORD")
}
func EnvSMTP_Host() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("SMTP_HOST")
}
