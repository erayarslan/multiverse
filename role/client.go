package role

import (
	"context"
	"log"
	"multipass-cluster/api"
	"multipass-cluster/config"
)

type client struct {
	apiClient api.Client
	cfg       *config.Config
}

func (c *client) Execute() error {
	log.Printf("api server addr: %s", c.cfg.APIServerAddr)

	var err error
	c.apiClient, err = api.NewClient(c.cfg.APIServerAddr)
	if err != nil {
		log.Fatalf("error while creating api client: %v", err)
	}

	names, err := c.apiClient.List(context.Background())
	if err != nil {
		log.Fatalf("error while listing: %v", err)
	}

	log.Printf("names: %v", names)

	return nil
}

func (c *client) GracefulShutdown() error {
	return c.apiClient.Close()
}

func NewClient(cfg *config.Config) Role {
	return &client{
		cfg: cfg,
	}
}
