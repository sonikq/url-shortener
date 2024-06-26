package app

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config конфигурация сервиса
type Config struct {
	HTTP HTTPConfig

	BaseURL string

	CtxTimeout int

	FileStoragePath string
	DatabaseDSN     string
	DBPoolWorkers   int

	LogLevel    string
	ServiceName string
}

// HTTPConfig -
type HTTPConfig struct {
	ServerAddress string
	Host          string
	Port          string
}

// Load -
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

	cfg.FileStoragePath = cast.ToString(os.Getenv("FILE_STORAGE_PATH"))

	cfg.DBPoolWorkers = cast.ToInt(os.Getenv("DB_POOL_WORKERS"))

	cfg.LogLevel = cast.ToString(os.Getenv("LOG_LEVEL"))
	cfg.ServiceName = cast.ToString(os.Getenv("SERVICE_NAME"))

	return cfg, nil

}

const (
	defaultServerAddress   = "localhost:8080"
	defaultBaseURL         = "http://localhost:8080"
	defaultLogLevel        = "info"
	defaultServiceName     = "url-shortener"
	defaultFileStoragePath = "/tmp/short-url-storage.json"
	defaultDatabaseDSN     = ""
	defaultDBPoolWorkers   = 250
)

// ParseConfig -
func ParseConfig(cfg *Config) {
	serverAddress := flag.String("a", defaultServerAddress, "server address defines on what port and host the server will be started")
	baseResURL := flag.String("b", defaultBaseURL, "defines which base address will be of resulting shortened URL")
	fileStoragePath := flag.String("f", defaultFileStoragePath, "determines where the data will be saved")
	databaseDSN := flag.String("d", defaultDatabaseDSN, "defines the database connection address")
	dbPoolWorkers := flag.Int("p", defaultDBPoolWorkers, "defines count of pool workers for db")
	flag.Parse()

	cfg.HTTP.ServerAddress = getEnvString("SERVER_ADDRESS", serverAddress)
	cfg.HTTP.Host = strings.Split(cfg.HTTP.ServerAddress, ":")[0]
	cfg.HTTP.Port = strings.Split(cfg.HTTP.ServerAddress, ":")[1]

	log.Printf("Server listening on %s port", cfg.HTTP.Port)

	cfg.BaseURL = getEnvString("BASE_URL", baseResURL)

	cfg.FileStoragePath = getEnvString("FILE_STORAGE_PATH", fileStoragePath)
	cfg.DatabaseDSN = getEnvString("DATABASE_DSN", databaseDSN)
	cfg.DBPoolWorkers = getEnvInt("DB_POOL_WORKERS", dbPoolWorkers)

	cfg.LogLevel = defaultLogLevel
	cfg.ServiceName = defaultServiceName
}

func getEnvString(key string, argumentValue *string) string {
	envValue, exists := os.LookupEnv(key)
	if !exists {
		return *argumentValue
	}
	return envValue
}

func getEnvInt(key string, argumentValue *int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err == nil {
		return value
	}
	return *argumentValue
}
