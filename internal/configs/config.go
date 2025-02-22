package configs

import (
	"flag"
	"os"
)

var (
	FlagRunAddr  string
	FlagBaseAddr string
	ShortURLLen  = 8
	LoggerLevel  = "info"
)

func ParseFlags() {
	flag.StringVar(&FlagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&FlagBaseAddr, "b", "http://localhost:8080", "basic address and port of result url")
	flag.Parse()
}

func ParseEnv() {
	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		FlagRunAddr = envRunAddr
	}
	if envRunAddr := os.Getenv("BASE_URL"); envRunAddr != "" {
		FlagBaseAddr = envRunAddr
	}
}
