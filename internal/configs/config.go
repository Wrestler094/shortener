// Package configs предоставляет функциональность для управления конфигурацией приложения.
// Поддерживает загрузку конфигурации из различных источников:
// - Флаги командной строки
// - JSON-файл конфигурации
// - Переменные окружения
package configs

import (
	"encoding/json"
	"flag"
	"os"
)

// Глобальные переменные для хранения флагов конфигурации
var (
	// FlagRunAddr - адрес и порт для запуска сервера (например, ":8080")
	FlagRunAddr string
	// FlagBaseAddr - базовый адрес и порт для результирующих URL (например, "http://localhost:8080")
	FlagBaseAddr string
	// FlagFileStoragePath - путь к файлу для сохранения настроек (например, "/path/to/storage.json")
	FlagFileStoragePath string
	// FlagDatabaseDSN - строка подключения к базе данных (например, "postgres://user:pass@localhost:5432/dbname")
	FlagDatabaseDSN string
	// FlagEnableHTTPS - флаг для включения HTTPS протокола
	FlagEnableHTTPS bool
	// flagConfigPath - путь к файлу конфигурации
	flagConfigPath string
)

// Константы приложения
var (
	// ShortURLLen - длина сокращенного URL в символах
	ShortURLLen = 8
	// LoggerLevel - уровень логирования (info, debug, warn, error)
	LoggerLevel = "info"
	// HTTPSPort - стандартный порт для HTTPS соединений
	HTTPSPort = 8443
)

// Config представляет структуру конфигурации приложения
type Config struct {
	// RunAddr - адрес и порт для запуска сервера
	RunAddr string `json:"server_address"`
	// BaseAddr - базовый адрес и порт для результирующих URL
	BaseAddr string `json:"base_url"`
	// FileStoragePath - путь к файлу для сохранения настроек
	FileStoragePath string `json:"file_storage_path"`
	// DatabaseDSN - строка подключения к базе данных
	DatabaseDSN string `json:"database_dsn"`
	// EnableHTTPS - флаг для включения HTTPS
	EnableHTTPS bool `json:"enable_https"`
}

// InitConfig инициализирует конфигурацию приложения.
// Последовательность загрузки конфигурации:
// 1. Флаги командной строки
// 2. JSON-файл конфигурации
// 3. Переменные окружения
func InitConfig() {
	ParseFlags()
	ParseJSON()
	ParseEnv()
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
	flag.BoolVar(&FlagEnableHTTPS, "s", false, "enable HTTPS protocol")
	flag.StringVar(&flagConfigPath, "config", "", "path to config file")
	flag.StringVar(&flagConfigPath, "c", "", "path to config file (shorthand)")

	/* DEV */
	//flag.StringVar(&FlagFileStoragePath, "f", "internal/storage/urls.json", "path to the file where current settings are saved")
	//flag.StringVar(&FlagDatabaseDSN, "d", "postgres://admin:secret@localhost:5432/mydatabase?sslmode=disable", "database connection")

	flag.Parse()
}

// ParseJSON загружает конфигурацию из JSON-файла.
// Приоритет значений:
// 1. Значения из флагов командной строки (если установлены)
// 2. Значения из JSON-файла (если флаг не установлен)
// Путь к файлу конфигурации может быть указан через:
// - Флаг командной строки -c или --config
// - Переменную окружения CONFIG
func ParseJSON() {
	if flagConfigPath == "" {
		flagConfigPath = os.Getenv("CONFIG")
	}

	if flagConfigPath != "" {
		fileCfg, err := LoadConfigFromFile(flagConfigPath)
		if err == nil {
			if flagIsSet("a") {
				FlagRunAddr = fileCfg.RunAddr
			}
			if flagIsSet("b") {
				FlagBaseAddr = fileCfg.BaseAddr
			}
			if FlagFileStoragePath == "" {
				FlagFileStoragePath = fileCfg.FileStoragePath
			}
			if FlagDatabaseDSN == "" {
				FlagDatabaseDSN = fileCfg.DatabaseDSN
			}
			if flagIsSet("s") {
				FlagEnableHTTPS = fileCfg.EnableHTTPS
			}
		}
	}
}

// ParseEnv загружает конфигурацию из переменных окружения.
// Поддерживаемые переменные окружения:
// SERVER_ADDRESS - адрес и порт для запуска сервера
// BASE_URL - базовый адрес и порт для результирующих URL
// FILE_STORAGE_PATH - путь к файлу для сохранения настроек
// DATABASE_DSN - строка подключения к базе данных
// ENABLE_HTTPS - флаг для включения HTTPS (значения: "true" или "false")
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
		if envRunAddr == "true" {
			FlagEnableHTTPS = true
		} else {
			FlagEnableHTTPS = false
		}
	}
}

// flagIsSet проверяет, был ли установлен флаг командной строки
// name - имя флага для проверки
// Возвращает true, если флаг был установлен
func flagIsSet(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// LoadConfigFromFile загружает конфигурацию из JSON-файла
// path - путь к файлу конфигурации
// Возвращает структуру Config и ошибку, если она возникла
func LoadConfigFromFile(path string) (Config, error) {
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}
