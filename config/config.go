package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)


type Configuration struct {
	Port       string `env:"PORT" envDefault:"8080"`
	SigningKey string `env:"SIGNING_KEY"`
	JwtSecret  string `env:"JWT_SECRET"`
	DBConnUrl  string `env:"DB_CONN_URL"`
}

func NewConfig(filename ...string) *Configuration{
	err := godotenv.Load(filename...)
	if err != nil {
		log.Printf("No .env found %q \n",filename)
	}
	cfg := Configuration{}
	err = env.Parse(&cfg);
	if err != nil {
		fmt.Printf("%v\n",err)
	}
	return &cfg
}

