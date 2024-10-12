package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	PackSizes []int
}

// LoadConfig loads the configuration from the .env file
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	packSizesStr := os.Getenv("PACK_SIZES")
	if packSizesStr == "" {
		return nil, fmt.Errorf("PACK_SIZES environment variable is not set")
	}

	packSizesStrSlice := strings.Split(packSizesStr, ",")
	packSizes := make([]int, len(packSizesStrSlice))

	for i, sizeStr := range packSizesStrSlice {
		size, err := strconv.Atoi(strings.TrimSpace(sizeStr))
		if err != nil {
			return nil, fmt.Errorf("invalid pack size: %s", sizeStr)
		}
		packSizes[i] = size
	}

	return &Config{
		PackSizes: packSizes,
	}, nil
}
