package main

import (
	"log"
	"multipass-cluster/config"
	"multipass-cluster/role"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	defer log.Printf("multipass cluster is shutting down")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cfg := config.NewConfig()

	closeCh := make(chan os.Signal, 1)
	doneCh := make(chan struct{}, 1)
	signal.Notify(closeCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT, syscall.SIGQUIT)

	var roles []role.Role

	if cfg.IsMaster {
		roles = append(roles, role.NewMaster(cfg))
	}

	if cfg.IsWorker {
		roles = append(roles, role.NewWorker(cfg))
	}

	if cfg.IsClient {
		roles = append(roles, role.NewClient(cfg, doneCh))
	}

	if len(roles) == 0 {
		log.Printf("no role selected")
		return
	}

	defer func() {
		for _, r := range roles {
			if err := r.GracefulShutdown(); err != nil {
				log.Fatalf("error while graceful shutdown: %v", err)
			}
		}
	}()

	for _, r := range roles {
		if err := r.Execute(); err != nil {
			log.Printf("error while executing role: %v", err)
			return
		}
	}

	select {
	case <-doneCh:
	case <-closeCh:
	}
}
