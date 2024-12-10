package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Client         *http.Client
	DefaultHeaders map[string]string
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func configSetup() *Config {
	// Set headers that try to be as close as possible to a real browser
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
	headers["Connection"] = "keep-alive"
	headers["Accept-Encoding"] = "gzip, deflate, br, zstd"
	headers["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	headers["Upgrade-Insecure-Request"] = "1"
	headers["Sec-Fetch-Dest"] = "document"
	headers["Sec-Fetch-Mode"] = "navigate"
	headers["Sec-Fetch-Site"] = "cross-site"
	headers["Sec-Fetch-User"] = "?1"
	headers["Priority"] = "u=0, i"
	headers["Accept-Language"] = "en-Us,en;q=0.9"
	headers["cache-control"] = "max-age=0"

	client := http.DefaultClient

	return &Config{
		Client:         client,
		DefaultHeaders: headers,
	}
}
