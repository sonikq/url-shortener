package app

import (
	"encoding/json"
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

	BaseURL string `json:"base_url"`

	CtxTimeout int

	FileStoragePath string `json:"file_storage_path"`
	DatabaseDSN     string `json:"database_dsn"`
	DBPoolWorkers   int

	TrustedSubnet string `json:"trusted_subnet"`
	UseGRPC       bool

	ConfigPath  string
	LogLevel    string
	ServiceName string
}

// HTTPConfig -
type HTTPConfig struct {
	ServerAddress string `json:"server_address"`
	Host          string
	Port          string
	EnableHTTPS   string `json:"enable_https"`
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
	cfg.ConfigPath = cast.ToString(os.Getenv("CONFIG"))
	cfg.TrustedSubnet = cast.ToString(os.Getenv("TRUSTED_SUBNET"))
	cfg.UseGRPC = cast.ToBool(os.Getenv("USE_GRPC"))

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
	defaultTLSRequire      = ""
	defaultConfigPath      = ""
	defaultTrustedSubnet   = ""
	defaultUseGRPC         = false
)

// ParseConfig -
func ParseConfig(cfg *Config) {
	serverAddress := flag.String("a", defaultServerAddress, "server address defines on what port and host the server will be started")
	baseResURL := flag.String("b", defaultBaseURL, "defines which base address will be of resulting shortened URL")
	fileStoragePath := flag.String("f", defaultFileStoragePath, "determines where the data will be saved")
	databaseDSN := flag.String("d", defaultDatabaseDSN, "defines the database connection address")
	dbPoolWorkers := flag.Int("p", defaultDBPoolWorkers, "defines count of pool workers for db")
	tlsRequire := flag.String("s", defaultTLSRequire, "server would be run on TLS")
	configPath := flag.String("c", defaultConfigPath, "path to config file")
	configPath = flag.String("config", *configPath, "path to config file")
	trustedSubnet := flag.String("t", defaultTrustedSubnet, "trusted subnetwork")
	useGRPC := flag.Bool("grpc", defaultUseGRPC, "whether to use grpc")
	flag.Parse()

	cfg.ConfigPath = getEnvString("CONFIG", configPath)
	if cfg.ConfigPath != "" {
		cfgFromFile, err := loadConfigFromFile(cfg.ConfigPath)
		if err != nil {
			log.Fatalf("cant parse config from file: %s", cfg.ConfigPath)
		}
		cfg = cfgFromFile
	}

	cfg.TrustedSubnet = getEnvString("TRUSTED_SUBNET", trustedSubnet)
	cfg.UseGRPC = getEnvBool("USE_GRPC", useGRPC)

	cfg.HTTP.ServerAddress = getEnvString("SERVER_ADDRESS", serverAddress)
	cfg.HTTP.Host = strings.Split(cfg.HTTP.ServerAddress, ":")[0]
	cfg.HTTP.Port = strings.Split(cfg.HTTP.ServerAddress, ":")[1]

	log.Printf("Server listening on %s port", cfg.HTTP.Port)

	cfg.BaseURL = getEnvString("BASE_URL", baseResURL)

	cfg.FileStoragePath = getEnvString("FILE_STORAGE_PATH", fileStoragePath)
	cfg.DatabaseDSN = getEnvString("DATABASE_DSN", databaseDSN)
	cfg.DBPoolWorkers = getEnvInt("DB_POOL_WORKERS", dbPoolWorkers)
	cfg.HTTP.EnableHTTPS = getEnvString("ENABLE_HTTPS", tlsRequire)
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

func getEnvBool(key string, argumentValue *bool) bool {
	value, err := strconv.ParseBool(os.Getenv(key))
	if err == nil {
		return value
	}
	return *argumentValue
}

func loadConfigFromFile(configPath string) (*Config, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	fileConfig := Config{
		HTTP: HTTPConfig{
			ServerAddress: defaultServerAddress,
			EnableHTTPS:   defaultTLSRequire,
		},
		BaseURL:         defaultBaseURL,
		CtxTimeout:      500,
		FileStoragePath: defaultFileStoragePath,
		DatabaseDSN:     defaultDatabaseDSN,
		DBPoolWorkers:   defaultDBPoolWorkers,
		ConfigPath:      defaultConfigPath,
		LogLevel:        defaultLogLevel,
		ServiceName:     defaultServiceName,
	}

	if err = json.NewDecoder(f).Decode(&fileConfig); err != nil {
		return nil, err
	}

	return &fileConfig, nil
}
