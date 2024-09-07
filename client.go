package go_alist

import (
	"net/http"
	"strings"
)

type Client struct {
	config     *Config
	httpClient *http.Client
}
type Config struct {
	ServerUrl string
	Token     string
}

func NewClient(config *Config) *Client {
	config.ServerUrl = strings.TrimSuffix(config.ServerUrl, "/")
	return &Client{
		config:     config,
		httpClient: &http.Client{},
	}
}
