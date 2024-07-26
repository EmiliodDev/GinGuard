package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
    Server struct {
        Port            int
    }
    Database struct {
        Driver          string
        Host            string
        Port            int
        User            string
        Password        string
        Name            string
    }
    JWT struct {
        SecretKey       string
        ExpireMinutes   int
    }
}

var AppConfig Config

func LoadConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %v", err)
    }

    if err := viper.Unmarshal(&AppConfig); err != nil {
        log.Fatalf("Unable to decode into a struct, %v", err)
    }
}


