package app

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	HTTP HTTPConfig

	CtxTimeout int

	LogLevel string
}

type HTTPConfig struct {
	Host string
	Port string
}

func Load(envFiles ...string) (Config, error) {

	if len(envFiles) != 0 {
		if err := godotenv.Load(envFiles...); err != nil {
			return Config{}, err
		}
	}

	var cfg = Config{}

	cfg.HTTP.Host = cast.ToString(os.Getenv("HTTP_HOST"))
	cfg.HTTP.Port = cast.ToString(os.Getenv("HTTP_PORT"))

	cfg.CtxTimeout = cast.ToInt(os.Getenv("CTX_TIMEOUT"))

	cfg.LogLevel = cast.ToString(os.Getenv("LOG_LEVEL"))

	return cfg, nil

}
