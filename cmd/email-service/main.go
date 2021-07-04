package main

import (
	"github.com/joho/godotenv"
	"github.com/quadgod/email-service-go/internal/app/server"
)

func main() {
	godotenv.Load()
	server.Start()
}
