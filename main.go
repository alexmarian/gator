package main

import (
	"fmt"
	"github.com/alexmarian/gator/internal/config"
	"log"
)

const configUserName = "lane"

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	cfg.SetUser(configUserName)
	fmt.Printf("Read config: %+v\n", cfg)

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
