package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConf struct {
	URL   string
	Token string
}

func (c *DBConf) GetURL() (string, error) {
	url := c.URL
	if len(c.Token) > 0 {
		url = url + "?authToken=" + c.Token
	}

	if len(url) == 0 {
		return "", errors.New("No DB url given.")
	}
	return url, nil
}

type Settings struct {
	DB       DBConf
	OrgName  string
	ApiToken string
}

func LoadSettings() Settings {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Fprintf(os.Stdout, ".env not found: %s\n", err)
	}

	url := os.Getenv("DB_URL")
	token := os.Getenv("DB_TOKEN")
	orgName := os.Getenv("ORG_NAME")
	apiToken := os.Getenv("API_TOKEN")

	return Settings{DBConf{url, token}, orgName, apiToken}
}
