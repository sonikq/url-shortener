package app

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"log"
	"os"
	"strings"
)

type Config struct {
	HTTP HTTPConfig

	BaseURL string

	CtxTimeout int

	LogLevel string
}

type HTTPConfig struct {
	ServerAddress string
	Host          string
	Port          string
}

func Load(envFiles ...string) (Config, error) {

	if len(envFiles) != 0 {
		if err := godotenv.Load(envFiles...); err != nil {
			return Config{}, err
		}
	}

	var cfg = Config{}

	cfg.HTTP.ServerAddress = cast.ToString(os.Getenv("SERVER_ADDRESS"))
	cfg.HTTP.Host = cast.ToString(os.Getenv("HTTP_HOST"))
	cfg.HTTP.Port = cast.ToString(os.Getenv("HTTP_PORT"))

	cfg.BaseURL = cast.ToString(os.Getenv("BASE_URL"))

	cfg.CtxTimeout = cast.ToInt(os.Getenv("CTX_TIMEOUT"))

	cfg.LogLevel = cast.ToString(os.Getenv("LOG_LEVEL"))

	return cfg, nil

}

const (
	defaultServerAddress = "localhost:8080"
	defaultBaseURL       = "http://localhost:8080/abcdef"
)

func ParseConfig(cfg *Config) {
	serverAddress := flag.String("a", defaultServerAddress, "server address defines on what port and host the server will be started")
	baseResURL := flag.String("b", defaultBaseURL, "defines which base address will be of resulting shortened URL")
	flag.Parse()

	cfg.HTTP.ServerAddress = getEnvString("SERVER_ADDRESS", *serverAddress)
	cfg.HTTP.Host = strings.Split(cfg.HTTP.ServerAddress, ":")[0]
	cfg.HTTP.Port = strings.Split(cfg.HTTP.ServerAddress, ":")[1]

	log.Printf("Server listening on %s port", cfg.HTTP.Port)

	cfg.BaseURL = getEnvString("BASE_URL", *baseResURL)
}

func getEnvString(key string, argumentValue string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return argumentValue
}
