package config

import (
	"flag"
	"log"
	"os"
	"runtime"
)

type Config struct {
	LaunchInstanceName    string
	MultipassKeyFilePath  string
	MultipassAddr         string
	MultipassProxyBind    string
	MultipassCertFilePath string
	MasterAddr            string
	ShellInstanceName     string
	APIServerAddr         string
	NodeName              string
	LaunchDiskSpace       string
	LaunchMemSize         string
	LaunchNumCores        string
	IsMaster              bool
	IsWorker              bool
	Shell                 bool
	Instances             bool
	Nodes                 bool
	Launch                bool
	Info                  bool
	IsClient              bool
}

func NewConfig() *Config {
	cfg := &Config{}

	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("error while getting user config dir: %v", err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("error while getting hostname: %v", err)
	}

	cfg.NodeName = hostname

	defaultMultipassAddr := "unix:///var/run/multipass_socket" // nolint:gosec
	if runtime.GOOS == "windows" {
		defaultMultipassAddr = "localhost:50051" // nolint:gosec
		dir += "/../Local"
	}
	if runtime.GOOS == "linux" {
		dir += "/../snap/multipass/current/data"
		defaultMultipassAddr = "unix:///var/snap/multipass/common/multipass_socket" // nolint:gosec
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
	flag.BoolVar(&cfg.Instances, "instances", false, "list instances")
	flag.BoolVar(&cfg.Nodes, "nodes", false, "list nodes")
	flag.BoolVar(&cfg.Shell, "shell", false, "run as shell")
	flag.BoolVar(&cfg.Launch, "launch", false, "launch instance")
	flag.BoolVar(&cfg.Info, "info", false, "get info")
	flag.StringVar(&cfg.ShellInstanceName, "shell-instance-name", "primary", "shell instance name")
	flag.StringVar(&cfg.LaunchInstanceName, "launch-instance-name", "primary", "launch instance name")
	flag.StringVar(&cfg.LaunchNumCores, "launch-num-cores", "1", "launch instance num cores")
	flag.StringVar(&cfg.LaunchMemSize, "launch-mem-size", "1G", "launch instance mem size")
	flag.StringVar(&cfg.LaunchDiskSpace, "launch-disk-space", "4G", "launch instance disk space")

	flag.Parse()

	return cfg
}
