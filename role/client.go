package role

import (
	"context"
	"log"
	"multiverse/api"
	"multiverse/config"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type client struct {
	apiClient api.Client
	cfg       *config.Config
	doneCh    chan struct{}
}

func (c *client) list() {
	list, err := c.apiClient.List(context.Background())
	if err != nil {
		log.Fatalf("error while list: %v", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Node Name", "Name"})
	rows := make([]table.Row, len(list.Instances))
	for i, n := range list.Instances {
		rows = append(rows, table.Row{i + 1, n.NodeName, n.InstanceName})
	}
	t.AppendRows(rows)
	t.Render()
}

func (c *client) shell() {
	err := c.apiClient.Shell(context.Background(), c.cfg.ShellNodeName, c.cfg.ShellInstanceName)
	if err != nil {
		log.Fatalf("error while shell: %v", err)
	}
}

func (c *client) Execute() error {
	log.Printf("api server addr: %s", c.cfg.APIServerAddr)

	var err error
	c.apiClient, err = api.NewClient(c.cfg.APIServerAddr)
	if err != nil {
		log.Fatalf("error while creating api client: %v", err)
	}

	if c.cfg.List {
		c.list()
	} else if c.cfg.Shell {
		c.shell()
	}

	c.doneCh <- struct{}{}

	return nil
}

func (c *client) GracefulShutdown() error {
	return c.apiClient.Close()
}

func NewClient(cfg *config.Config, doneCh chan struct{}) Role {
	return &client{
		cfg:    cfg,
		doneCh: doneCh,
	}
}
