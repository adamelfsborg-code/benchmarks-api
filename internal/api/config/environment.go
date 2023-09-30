package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort uint16

	DBAddr     string
	DBUser     string
	DBPassword string
	DBDatabase string
}

func Build() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, fmt.Errorf("Failed to load Config: %w", err)
	}

	const base = 10
	const bitSize = 64

	config := Config{}

	serverPort, exists := os.LookupEnv("SERVER_PORT")
	if exists == false {
		return Config{}, fmt.Errorf("Failed to load Config: %w", err)
	}
	dbAddr, exists := os.LookupEnv("DB_ADDR")
	if exists == false {
		return Config{}, fmt.Errorf("Failed to load Config: %w", err)
	}
	dbUser, exists := os.LookupEnv("DB_USER")
	if exists == false {
		return Config{}, fmt.Errorf("Failed to load Config: %w", err)
	}
	dbPassword, exists := os.LookupEnv("DB_PASSWORD")
	if exists == false {
		return Config{}, fmt.Errorf("Failed to load Config: %w", err)
	}
	dbDatabase, exists := os.LookupEnv("DB_DATABASE")
	if exists == false {
		return Config{}, fmt.Errorf("Failed to load Config: %w", err)
	}

	port, err := strconv.ParseInt(serverPort, base, bitSize)
	if err == nil {
		config.ServerPort = uint16(port)
	}

	config.DBAddr = dbAddr
	config.DBUser = dbUser
	config.DBPassword = dbPassword
	config.DBDatabase = dbDatabase

	return config, nil
}
