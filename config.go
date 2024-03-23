package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Port             string
	ConnectionString string
}

func LoadConfiguration() (*Configuration, error) {
	godotenv.Load()
	port := os.Getenv("PORT")
	connStr := os.Getenv("CONN")

	err := validateNotEmpty(port, connStr)
	if err != nil {
		return nil, err
	}

	return &Configuration{Port: port, ConnectionString: connStr}, nil
}

func validateNotEmpty(params ...string) error {
	for _, p := range params {
		if p == "" {
			return fmt.Errorf("required env var empty")
		}
	}

	return nil
}
