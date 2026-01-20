package main

import (
	"gw-exchanger/internal/config"
	"log"
)

func main() {
	cfg := config.GetConfig()

	log.Printf("Config: %+v", cfg)
}
