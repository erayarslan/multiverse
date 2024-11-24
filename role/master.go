package role

import (
	"log"

	"github.com/erayarslan/multiverse/api"
	"github.com/erayarslan/multiverse/cluster"
	"github.com/erayarslan/multiverse/config"
)

type master struct {
	cfg *config.Config
}

func (c *master) Execute() error {
	log.Printf("master addr: %s", c.cfg.MasterAddr)

	clusterServer, err := cluster.NewServer(c.cfg.MasterAddr)
	if err != nil {
		log.Fatalf("error while creating master: %v", err)
	}

	log.Printf("api server addr: %s", c.cfg.APIServerAddr)

	apiServer, err := api.NewServer(c.cfg.APIServerAddr, clusterServer)
	if err != nil {
		log.Fatalf("error while creating api server: %v", err)
	}

	go func() {
		var err error
		if err = clusterServer.Serve(); err != nil {
			log.Fatalf("error while serving api server: %v", err)
		}
	}()

	go func() {
		var err error
		if err = apiServer.Serve(); err != nil {
			log.Fatalf("error while serving master: %v", err)
		}
	}()

	return nil
}

func (c *master) GracefulShutdown() error {
	return nil
}

func NewMaster(cfg *config.Config) Role {
	return &master{
		cfg: cfg,
	}
}
