package role

import (
	"log"
	"multiverse/agent"
	"multiverse/cluster"
	"multiverse/config"
	"multiverse/multipass"
)

type worker struct {
	cfg           *config.Config
	clusterClient cluster.Client
}

func (c *worker) Execute() error {
	log.Printf("master to connect addr: %s", c.cfg.MasterAddr)

	multipassClient, err := multipass.NewClient(
		c.cfg.MultipassAddr,
		c.cfg.MultipassCertFilePath,
		c.cfg.MultipassKeyFilePath,
	)
	if err != nil {
		log.Fatalf("error while creating multipass client: %v", err)
	}

	state := agent.NewState(multipassClient)
	go state.Run()

	server, err := agent.NewServer(c.cfg.MultipassProxyBind, multipassClient)
	if err != nil {
		log.Fatalf("error while creating multipass proxy: %v", err)
	}

	c.clusterClient, err = cluster.NewClient(c.cfg.MasterAddr, c.cfg.NodeName, server, multipassClient, state)
	if err != nil {
		log.Fatalf("error while creating worker: %v", err)
	}

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("error while serving multipass proxy: %v", err)
		}
	}()

	go func() {
		err = c.clusterClient.Sync()
		if err != nil {
			log.Fatalf("error while sync to master: %v", err)
		}
	}()

	return nil
}

func (c *worker) GracefulShutdown() error {
	return c.clusterClient.Close()
}

func NewWorker(cfg *config.Config) Role {
	return &worker{
		cfg: cfg,
	}
}
