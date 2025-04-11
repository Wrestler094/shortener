package configs

import (
	"flag"
	"os"
)

/* FLAGS */
var (
	FlagRunAddr         string
	FlagBaseAddr        string
	FlagFileStoragePath string
	FlagDatabaseDSN     string
)

/* CONSTANTS */
var (
	ShortURLLen = 8
	LoggerLevel = "info"
)

func ParseFlags() {
	flag.StringVar(&FlagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&FlagBaseAddr, "b", "http://localhost:8080", "basic address and port of result url")
	flag.StringVar(&FlagFileStoragePath, "f", "internal/storage/urls.json", "path to the file where current settings are saved")
	flag.StringVar(&FlagDatabaseDSN, "d", "", "database connection")
	flag.Parse()
}

func ParseEnv() {
	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		FlagRunAddr = envRunAddr
	}
	if envRunAddr := os.Getenv("BASE_URL"); envRunAddr != "" {
		FlagBaseAddr = envRunAddr
	}
	if envRunAddr := os.Getenv("FILE_STORAGE_PATH"); envRunAddr != "" {
		FlagFileStoragePath = envRunAddr
	}
	if envRunAddr := os.Getenv("DATABASE_DSN"); envRunAddr != "" {
		FlagDatabaseDSN = envRunAddr
	}
}
