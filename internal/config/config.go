package config

import (
	"log"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
)

type Config struct {
	LogLevel   string
	HTTPPort   string
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

func NewConfig() *Config {
	var cfg Config
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipFlags: true,
		Files:     []string{".env"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".env": aconfigdotenv.New(),
		},
	})
	if err := loader.Load(); err != nil {
		log.Fatal(err)
	}
	return &cfg
}
