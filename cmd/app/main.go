package main

import (
	"fmt"

	"github.com/SunilKividor/PillNet-Backend/internal/di"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err == nil {
		fmt.Println("[INFO] Loaded .env from project root")
	}

	server, err := di.InitializeApp()
	if err != nil {
		panic(err)
	}

	err = server.Serve()
	if err != nil {
		panic(err)
	}
}
