package android

import (
	"log"
	"time"

	"mc-client/core"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Start(host, path string) error {
	cfg := core.Config{
		Host: host,
		Path: path,
	}
	log.Println("Connecting to", host+path)
	return core.RunBedrock(cfg)
}
