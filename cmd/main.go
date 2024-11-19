package main

import (
	"context"
	"flag"
	"log"
	"multipass-cluster/agent"
	"multipass-cluster/api"
	"multipass-cluster/cluster"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	isMaster              bool
	isWorker              bool
	isClient              bool
	masterAddr            string
	apiServerAddr         string
	multipassAddr         string
	multipassProxyBind    string
	multipassCertFilePath string
	multipassKeyFilePath  string
)

func main() {
	dir, err := os.UserConfigDir()
	if err != nil {
		return
	}

	defaultMultipassAddr := "unix:///var/run/multipass_socket"
	if runtime.GOOS == "windows" {
		defaultMultipassAddr = "localhost:50051"
		dir = dir + "/../Local"
	}

	defaultMultipassCertFilePath := dir + "/multipass-client-certificate/multipass_cert.pem"
	defaultMultiPassKeyFilePath := dir + "/multipass-client-certificate/multipass_cert_key.pem"

	flag.BoolVar(&isMaster, "master", false, "run as master")
	flag.BoolVar(&isWorker, "worker", false, "run as worker")
	flag.BoolVar(&isClient, "client", false, "run as client")

	flag.StringVar(&masterAddr, "master-addr", "localhost:1337", "master addr to listen on")
	flag.StringVar(&apiServerAddr, "api-server-addr", "localhost:1338", "api server addr to listen on")
	flag.StringVar(&multipassAddr, "multipass-addr", defaultMultipassAddr, "multipass addr to connect")
	flag.StringVar(&multipassProxyBind, "multipass-proxy-bind", "localhost", "multipass proxy bind to listen on")

	flag.StringVar(&multipassCertFilePath, "multipass-cert-file", defaultMultipassCertFilePath, "multipass cert file for tls")
	flag.StringVar(&multipassKeyFilePath, "multipass-key-file", defaultMultiPassKeyFilePath, "multipass key file for tls")

	flag.Parse()

	doneCh := make(chan struct{})
	closeCh := make(chan os.Signal, 1)
	signal.Notify(closeCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGABRT, syscall.SIGQUIT)

	var clusterServerCh = make(chan cluster.Server, 1)

	go func() {
		if !isMaster {
			return
		}

		log.Printf("master addr: %s", masterAddr)

		err := cluster.NewServer(masterAddr, clusterServerCh)
		if err != nil {
			log.Fatalf("error while creating master: %v", err)
		}
	}()

	go func() {
		if !isMaster {
			return
		}

		log.Printf("api server addr: %s", apiServerAddr)

		err = api.NewServer(apiServerAddr, clusterServerCh)
		if err != nil {
			log.Fatalf("error while creating api server: %v", err)
		}
	}()

	var agentInfoCh = make(chan *agent.Info, 1)

	go func() {
		if !isWorker {
			return
		}

		err := agent.NewServer(multipassAddr, multipassProxyBind, multipassCertFilePath, multipassKeyFilePath, agentInfoCh)
		if err != nil {
			log.Fatalf("error while creating multipass proxy: %v", err)
		}
	}()

	go func() {
		if !isWorker {
			return
		}

		log.Printf("master addr: %s", masterAddr)

		client, err := cluster.NewClient(masterAddr, agentInfoCh)
		if err != nil {
			log.Fatalf("error while creating worker: %v", err)
		}
		defer func(client cluster.Client) {
			err := client.Close()
			if err != nil {
				log.Fatalf("error while closing worker: %v", err)
			}
		}(client)

		err = client.Join()
		if err != nil {
			log.Fatalf("error while joining master: %v", err)
		}
	}()

	if isClient {
		log.Printf("api server addr: %s", apiServerAddr)

		client, err := api.NewClient(apiServerAddr)
		if err != nil {
			log.Fatalf("error while creating api client: %v", err)
		}
		defer func(client api.Client) {
			err := client.Close()
			if err != nil {
				log.Fatalf("error while closing api client: %v", err)
			}
		}(client)

		names, err := client.List(context.Background())
		if err != nil {
			log.Fatalf("error while listing: %v", err)
		}

		log.Printf("names: %v", names)

		return
	}

	select {
	case <-doneCh:
		log.Printf("done")
	case <-closeCh:
		log.Printf("shutting down...")
	}
}
