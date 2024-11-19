package config

import (
	"flag"
	"log"
	"os"
	"runtime"
)

type Config struct {
	MasterAddr            string
	APIServerAddr         string
	MultipassAddr         string
	MultipassProxyBind    string
	MultipassCertFilePath string
	MultipassKeyFilePath  string
	IsMaster              bool
	IsWorker              bool
	IsClient              bool
}

func NewConfig() *Config {
	cfg := &Config{}

	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("error while getting user config dir: %v", err)
	}

	defaultMultipassAddr := "unix:///var/run/multipass_socket" // nolint:gosec
	if runtime.GOOS == "windows" {
		defaultMultipassAddr = "localhost:50051" // nolint:gosec
		dir += "/../Local"
	}

	defaultMultipassCertFilePath := dir + "/multipass-client-certificate/multipass_cert.pem"
	defaultMultiPassKeyFilePath := dir + "/multipass-client-certificate/multipass_cert_key.pem"

	flag.BoolVar(&cfg.IsMaster, "master", false, "run as master")
	flag.BoolVar(&cfg.IsWorker, "worker", false, "run as worker")
	flag.BoolVar(&cfg.IsClient, "client", false, "run as client")
	flag.StringVar(&cfg.MasterAddr, "master-addr", "localhost:1337", "master addr to listen on")
	flag.StringVar(&cfg.APIServerAddr, "api-server-addr", "localhost:1338", "api server addr to listen on")
	flag.StringVar(&cfg.MultipassAddr, "multipass-addr", defaultMultipassAddr, "multipass addr to connect")
	flag.StringVar(&cfg.MultipassProxyBind, "multipass-proxy-bind", "localhost", "multipass proxy bind to listen on")
	flag.StringVar(&cfg.MultipassCertFilePath, "multipass-cert-file", defaultMultipassCertFilePath, "multipass cert file for tls")
	flag.StringVar(&cfg.MultipassKeyFilePath, "multipass-key-file", defaultMultiPassKeyFilePath, "multipass key file for tls")

	flag.Parse()

	return cfg
}
