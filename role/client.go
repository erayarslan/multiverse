package role

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/erayarslan/multiverse/common"

	"github.com/erayarslan/multiverse/api"
	"github.com/erayarslan/multiverse/config"
)

type client struct {
	apiClient api.Client
	cfg       *config.Config
	doneCh    chan struct{}
}

func (c *client) instances() {
	getInstancesReply, err := c.apiClient.Instances(context.Background())
	if err != nil {
		log.Fatalf("error while instances: %v", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)

	fs := "%s\t%s\t%s\t%s\t%s\n"
	_, err = fmt.Fprintf(w, fs, "Node Name", "Instance Name", "State", "IPv4", "Image")
	if err != nil {
		return
	}
	for _, n := range getInstancesReply.Instances {
		_, err = fmt.Fprintf(w, fs, n.NodeName, n.Instance.Name, n.Instance.State, strings.Join(n.Instance.Ipv4, "\n"), n.Instance.Image)
		if err != nil {
			return
		}
	}

	err = w.Flush()
	if err != nil {
		return
	}
}

func (c *client) nodes() {
	getNodesReply, err := c.apiClient.Nodes(context.Background())
	if err != nil {
		log.Fatalf("error while nodes: %v", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)

	fs := "%s\t%s\t%s\t%s\t%s\t%s\n"
	_, err = fmt.Fprintf(w, fs, "Node Name", "IPv4", "Cpu", "Mem", "Disk", "Last Sync")
	if err != nil {
		return
	}
	for _, n := range getNodesReply.Nodes {
		_, err = fmt.Fprintf(w, fs,
			n.Name,
			strings.Join(n.Ipv4, "\n"),
			fmt.Sprintf("%d", n.Resource.Cpu.Available),
			fmt.Sprintf("%vGb", n.Resource.Memory.Available/1024/1024/1024),
			fmt.Sprintf("%vGb", n.Resource.Disk.Available/1024/1024/1024),
			n.LastSync.AsTime().Format("2006-01-02 15:04:05 MST"),
		)
		if err != nil {
			return
		}
	}

	err = w.Flush()
	if err != nil {
		return
	}
}

func (c *client) info() {
	getInfoReply, err := c.apiClient.Info(context.Background())
	if err != nil {
		log.Fatalf("error while info: %v", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)

	fs := "%s\t%s\t%s\t%s\t%s\t%s\n"
	_, err = fmt.Fprintf(w, fs, "Node Name", "Instance Name", "Cpu", "Load", "Disk", "Memory")
	if err != nil {
		return
	}
	for _, n := range getInfoReply.Instances {
		diskUsage, err := strconv.ParseFloat(n.Instance.DiskUsage, 64)
		if err != nil {
			log.Fatalf("error while parsing disk usage: %v", err)
		}
		diskTotal, err := strconv.ParseFloat(n.Instance.DiskTotal, 64)
		if err != nil {
			log.Fatalf("error while parsing disk total: %v", err)
		}
		memoryUsage, err := strconv.ParseFloat(n.Instance.MemoryUsage, 64)
		if err != nil {
			log.Fatalf("error while parsing memory usage: %v", err)
		}
		memoryTotal, err := strconv.ParseFloat(n.Instance.MemoryTotal, 64)
		if err != nil {
			log.Fatalf("error while parsing memory total: %v", err)
		}

		_, err = fmt.Fprintf(w, fs, n.NodeName, n.Instance.Name, n.Instance.CpuCount, n.Instance.Load,
			fmt.Sprintf("%.1fGiB out of %.1fGiB", diskUsage/1024/1024/1024, diskTotal/1024/1024/1024),
			fmt.Sprintf("%.1fGiB out of %.1fGiB", memoryUsage/1024/1024/1024, memoryTotal/1024/1024/1024),
		)
		if err != nil {
			return
		}
	}

	err = w.Flush()
	if err != nil {
		return
	}
}

func (c *client) shell() {
	err := c.apiClient.Shell(context.Background(), c.cfg.ShellInstanceName)
	if err != nil {
		log.Fatalf("error while shell: %v", err)
	}
}

func (c *client) launch() {
	// parse c.cfg.LaunchNumCores to int32
	numCores, err := strconv.ParseInt(c.cfg.LaunchNumCores, 10, 32)
	if err != nil {
		log.Fatalf("error while parsing num cores: %v", err)
	}

	_, err = c.apiClient.Launch(context.Background(), &common.LaunchRequest{
		InstanceName: c.cfg.LaunchInstanceName,
		NumCores:     int32(numCores),
		MemSize:      c.cfg.LaunchMemSize,
		DiskSpace:    c.cfg.LaunchDiskSpace,
	})
	if err != nil {
		log.Fatalf("error while launch: %v", err)
	}

	log.Printf("instance %s launched", c.cfg.LaunchInstanceName)
}

func (c *client) Execute() error {
	log.Printf("api server addr: %s", c.cfg.APIServerAddr)

	var err error
	c.apiClient, err = api.NewClient(c.cfg.APIServerAddr)
	if err != nil {
		log.Fatalf("error while creating api client: %v", err)
	}

	switch {
	case c.cfg.Instances:
		c.instances()
	case c.cfg.Nodes:
		c.nodes()
	case c.cfg.Shell:
		c.shell()
	case c.cfg.Launch:
		c.launch()
	case c.cfg.Info:
		c.info()
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
