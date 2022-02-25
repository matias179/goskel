package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	envPath := os.ExpandEnv("./dotenv/.env.${APP_ENV}")

	_, err := os.Stat("/.dockerenv")

	if err == nil || os.Getenv("ECS") != "" {
		envPath = os.ExpandEnv("/bin/dotenv/.env.${APP_ENV}")
	}

	err = godotenv.Load(envPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello World!")
}
