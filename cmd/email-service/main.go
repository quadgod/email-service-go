package main

import (
	"github.com/joho/godotenv"
	app "github.com/quadgod/email-service-go/internal/app"
)

func main() {
	godotenv.Load()
	app.Start()
}
