package agent

import (
	"context"
	"log"
	"multiverse/multipass"
	"sync"
	"time"
)

type state struct {
	multipassClient multipass.Client
	stateChan       chan state
	Instances       []*Instance
	stateMu         sync.RWMutex
}

type State interface {
	Listen() <-chan state
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

func (s *state) Run() {
	for {
		s.stateMu.Lock()
		s.updateInstances()
		s.stateChan <- *s
		s.stateMu.Unlock()
		time.Sleep(10 * time.Second)
	}
}

func (s *state) Listen() <-chan state {
	return s.stateChan
}

func NewState(multipassClient multipass.Client) State {
	return &state{
		multipassClient: multipassClient,
		stateMu:         sync.RWMutex{},
		stateChan:       make(chan state),
	}
}
