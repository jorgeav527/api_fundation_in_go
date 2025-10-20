package main

import (
	"fmt"
	"sync"
)

type Service struct {
	Name string
}

type Services struct {
	sync.RWMutex
	ServiceList map[string]Service
}

func InitServices() *Services {
	return &Services{
		// ServiceList: make(map[string]Service),
		ServiceList: map[string]Service{},
	}
}

func (s *Services) AddService(name string, srv Service) {
	s.Lock()
	defer s.Unlock()
	s.ServiceList[name] = srv
}

func (s *Services) GetOneService(name string) (Service, bool) {
	s.RLock()
	defer s.RUnlock()
	srv, ok := s.ServiceList[name]
	return srv, ok

}

func main() {
	s := InitServices()
	s.AddService("NoSQL", Service{"MongoDB"})
	srv, ok := s.GetOneService("NoSQL")
	if ok {
		fmt.Println("Found", srv.Name)
	}
}
