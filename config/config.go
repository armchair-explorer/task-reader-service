package config

import "os"

type Config struct {
    DBURL string
}

func LoadConfig() Config {
    return Config{
        DBURL: os.Getenv("DB_URL"),
    }
}

