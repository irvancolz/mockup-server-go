package main

import (
	"log"
	"mockup_server/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Println("failed to read env :", errEnv)
	}

	port := ":" + os.Getenv("ACTIVE_PORT")
	if err := router.GinRouter().Run(port); err != nil {
		log.Fatalln("failed to start server :", err)
	}
}
