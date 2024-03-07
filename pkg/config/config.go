package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	//.env := "workshop2-bcc.env"
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("error load env : %v", err)
	}
}
