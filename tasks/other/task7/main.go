// собес: Авито Платформа
// Задание:
// Реализовать универсальный конфиг, который бы использовал опциональный паттерн.

package main

import (
	"time"
)

type Config struct {
	Timeout time.Duration
	Retries int
	Debug   bool
}

type Option func(*Config)

func defaultConf() Config {
	return Config{
		Timeout: 5 * time.Second,
		Retries: 5,
		Debug:   false,
	}
}

func New(options ...Option) *Config {
	cfg := defaultConf()

	for _, option := range options {
		option(&cfg)
	}
	return &cfg
}

func WithTimeout(d time.Duration) Option {
	return func(c *Config) {
		c.Timeout = d
	}
}

func WithRetries(r int) Option {
	return func(c *Config) {
		c.Retries = r
	}
}

func WithDebug(d bool) Option {
	return func(c *Config) {
		c.Debug = d
	}
}
