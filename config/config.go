package config

import (
	"fmt"
	"os" 
	"strconv"
	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version string
	ServiceName string
	HttpPort int
}

// loadConfig loads configuration from environment variables
func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}
	
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HTTP Port is required")
	    os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Invalid HTTP Port")
		os.Exit(1)
	}

	configurations = Config{
		Version: 		version,
		ServiceName: 	serviceName,
		HttpPort: 		int(port),
	}
}


// GetConfig returns the loaded configuration
func GetConfig() Config {
	loadConfig()
	return configurations
}