package main

import (
	"log"

	"github.com/sergen-x/sergen-x-api/internal/server"
)

func main() {
	if err := server.Start(); err != nil {
		log.Fatalf("error while launching server: %v", err)
	}
}
