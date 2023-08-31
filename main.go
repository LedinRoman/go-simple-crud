package main

import (
	"fmt"
	"go-postgres/router"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load(".env")
    if err != nil {
		log.Fatalf("Error loading .env file")
    }

	r := router.Router()
	
	var port string = ":" + os.Getenv("APP_PORT")
	fmt.Println("Starting server on the port " + port)

	log.Fatal(http.ListenAndServe(port, r))
}
