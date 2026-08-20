package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type Config struct {
	HTTPAddr        string
	ShutdownTimeout time.Duration
	MaxWorkers      int
	DatabaseURL     string
	ObjectStoreURL  string
}

func Load() Config {
	c := Config{HTTPAddr: env("HTTP_ADDR", ":8088"), ShutdownTimeout: duration("SHUTDOWN_TIMEOUT", 10*time.Second), MaxWorkers: number("MAX_WORKERS", 8), DatabaseURL: os.Getenv("DATABASE_URL"), ObjectStoreURL: os.Getenv("OBJECT_STORE_URL")}
	return c
}
func Validate(c Config) error {
	if c.HTTPAddr == "" {
		return errors.New("http address required")
	}
	if c.MaxWorkers < 1 {
		return errors.New("max workers must be positive")
	}
	return nil
}
func env(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
func number(k string, d int) int {
	v, e := strconv.Atoi(os.Getenv(k))
	if e != nil || v < 1 {
		return d
	}
	return v
}
func duration(k string, d time.Duration) time.Duration {
	v, e := time.ParseDuration(os.Getenv(k))
	if e != nil {
		return d
	}
	return v
}
