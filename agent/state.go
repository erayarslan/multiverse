package agent

import (
	"context"
	"log"
	"math"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/disk"

	"github.com/erayarslan/multiverse/multipass"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type state struct {
	multipassClient multipass.Client
	stateChan       chan state
	Resource        *Resource
	Instances       []*Instance
	stateMu         sync.RWMutex
}

type State interface {
	Listen() <-chan state
	GetState() *state
	Run()
}

func (s *state) updateInstances() {
	res, err := s.multipassClient.List(context.Background())
	if err != nil {
		log.Printf("error while listing multipass: %v", err)
	} else {
		instances := make([]*Instance, 0, len(res))
		for _, instance := range res {
			instances = append(instances, &Instance{
				Name:  instance.Name,
				State: instance.State,
				Ipv4:  instance.Ipv4,
				Image: instance.Image,
			})
		}

		s.Instances = instances
	}
}

func (s *state) GetState() *state {
	s.stateMu.RLock()
	defer s.stateMu.RUnlock()
	return s
}

func (s *state) updateResources() {
	virtualMemoryStat, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("error while getting virtual memory: %v", err)
		return
	}
	cpuInfoStats, err := cpu.Info()
	if err != nil {
		log.Printf("error while getting cpu info: %v", err)
		return
	}
	percents, err := cpu.Percent(0, false)
	if err != nil {
		log.Printf("error while getting cpu percent: %v", err)
		return
	}
	var totalCores int32
	for _, cpuInfo := range cpuInfoStats {
		totalCores += cpuInfo.Cores
	}
	availableCore := totalCores - int32(math.Ceil(float64(totalCores)*percents[0]/100))
	diskUsage, err := disk.Usage("/")
	if err != nil {
		log.Printf("error while getting disk usage: %v", err)
		return
	}

	s.Resource = &Resource{
		Cpu: &CPU{
			Total:     totalCores,
			Available: availableCore,
		},
		Memory: &Memory{
			Total:     virtualMemoryStat.Total,
			Available: virtualMemoryStat.Available,
		},
		Disk: &Disk{
			Total:     diskUsage.Total,
			Available: diskUsage.Free,
		},
	}
}

func (s *state) Run() {
	for {
		s.stateMu.Lock()
		s.updateInstances()
		s.updateResources()
		s.stateChan <- *s
		s.stateMu.Unlock()
		time.Sleep(10 * time.Second)
	}
}

func (s *state) Listen() <-chan state {
	return s.stateChan
}

func NewState(multipassClient multipass.Client) State {
	s := &state{
		multipassClient: multipassClient,
		stateMu:         sync.RWMutex{},
		stateChan:       make(chan state),
	}
	return s
}
