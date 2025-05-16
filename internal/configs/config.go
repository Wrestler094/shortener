package configs

import (
	"encoding/json"
	"flag"
	"os"
	"strconv"
)

// Глобальные переменные для хранения флагов конфигурации
var (
	// FlagRunAddr - адрес и порт для запуска сервера
	FlagRunAddr string
	// FlagBaseAddr - базовый адрес и порт для результирующих URL
	FlagBaseAddr string
	// FlagFileStoragePath - путь к файлу для сохранения настроек
	FlagFileStoragePath string
	// FlagDatabaseDSN - строка подключения к базе данных
	FlagDatabaseDSN string
	// FlagEnableHTTPS - флаг для включения HTTPS
	FlagEnableHTTPS string
)

// Константы приложения
var (
	// ShortURLLen - длина сокращенного URL
	ShortURLLen = 8
	// LoggerLevel - уровень логирования
	LoggerLevel = "info"
	// HTTPS порт
	HTTPSPort = 443
)

type Config struct {
	RunAddr         string `json:"server_address"`
	BaseAddr        string `json:"base_url"`
	FileStoragePath string `json:"file_storage_path"`
	DatabaseDSN     string `json:"database_dsn"`
	EnableHTTPS     bool   `json:"enable_https"`
}

func ParseJSON() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.StringVar(&configPath, "c", "", "path to config file (shorthand)")
	flag.Parse() // парсим предварительно, чтобы получить config-файл

	if configPath == "" {
		configPath = os.Getenv("CONFIG")
	}

	if configPath != "" {
		fileCfg, err := LoadConfigFromFile(configPath)
		if err == nil {
			FlagRunAddr = fileCfg.RunAddr
			FlagBaseAddr = fileCfg.BaseAddr
			FlagFileStoragePath = fileCfg.FileStoragePath
			FlagDatabaseDSN = fileCfg.DatabaseDSN
			FlagEnableHTTPS = strconv.FormatBool(fileCfg.EnableHTTPS)
		}
	}
}

// ParseFlags парсит флаги командной строки и устанавливает значения конфигурации.
// Поддерживаемые флаги:
// -a: адрес и порт для запуска сервера (по умолчанию ":8080")
// -b: базовый адрес и порт для результирующих URL (по умолчанию "http://localhost:8080")
// -f: путь к файлу для сохранения настроек
// -d: строка подключения к базе данных
// -s: флаг для включения HTTPS
func ParseFlags() {
	flag.StringVar(&FlagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&FlagBaseAddr, "b", "http://localhost:8080", "basic address and port of result url")
	flag.StringVar(&FlagFileStoragePath, "f", "", "path to the file where current settings are saved")
	flag.StringVar(&FlagDatabaseDSN, "d", "", "database connection")
	flag.StringVar(&FlagEnableHTTPS, "s", "", "enable HTTPS protocol")

	/* DEV */
	//flag.StringVar(&FlagFileStoragePath, "f", "internal/storage/urls.json", "path to the file where current settings are saved")
	//flag.StringVar(&FlagDatabaseDSN, "d", "postgres://admin:secret@localhost:5432/mydatabase?sslmode=disable", "database connection")

	flag.Parse()
}

// ParseEnv парсит переменные окружения и устанавливает значения конфигурации.
// Поддерживаемые переменные окружения:
// SERVER_ADDRESS: адрес и порт для запуска сервера
// BASE_URL: базовый адрес и порт для результирующих URL
// FILE_STORAGE_PATH: путь к файлу для сохранения настроек
// DATABASE_DSN: строка подключения к базе данных
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
	if envRunAddr := os.Getenv("ENABLE_HTTPS"); envRunAddr != "" {
		FlagEnableHTTPS = envRunAddr
	}
}

func LoadConfigFromFile(path string) (Config, error) {
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}
