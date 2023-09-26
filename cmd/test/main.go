package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load();
    if err != nil {
        log.Fatal("error opening .env file")
    }

    thread_count := os.Getenv("THREAD_COUNT")
    log.Println(thread_count)
}
