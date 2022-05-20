package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")

	router.Router()
	fmt.Println("Server started on port " + port)
	http.ListenAndServe(":"+port, nil)
}
