package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var configurations *Config

type DBConfig struct {
	Host          string
	Port          int
	User          string
	Name          string
	Password      string
	EnableSSLMode bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
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

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT Secret Key is required")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("Database Host is required")
		os.Exit(1)
	}

	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		fmt.Println("Database Port is required")
		os.Exit(1)
	}

	dbPort, err := strconv.ParseInt(dbPortStr, 10, 64)
	if err != nil {
		fmt.Println("Database Port must be a number")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("Database Name is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("Database User is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("Database Password is required")
		os.Exit(1)
	}

	enableSSLMode := os.Getenv("DB_ENABLE_SSL_MODE")

	enblSSLMode, err := strconv.ParseBool(enableSSLMode)
	if err != nil {
		fmt.Println("Invalid DB Enable SSL Mode")
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		Host:          dbHost,
		Port:          int(dbPort),
		Name:          dbName,
		User:          dbUser,
		Password:      dbPassword,
		EnableSSLMode: enblSSLMode,
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}
}

// GetConfig returns the loaded configuration
func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
