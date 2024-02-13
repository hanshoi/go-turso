package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	URL   string
	Token string
}

func (c *Settings) GetURL() (string, error) {
	url := c.URL
	if len(c.Token) > 0 {
		url = url + "?authToken=" + c.Token
	}

	if len(url) == 0 {
		return "", errors.New("No DB url given.")
	}
	return url, nil
}

func LoadSettings() Settings {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Fprintf(os.Stdout, ".env not found: %s\n", err)
	}

	url := os.Getenv("DB_URL")
	token := os.Getenv("DB_TOKEN")

	return Settings{url, token}
}
