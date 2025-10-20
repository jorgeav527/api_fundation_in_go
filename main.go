package main

import (
	"fmt"
	"runtime"
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

func (s *Services) runWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	serviveName := fmt.Sprintf("Service-%d", id)
	s.AddService(serviveName, Service{fmt.Sprintf("Worker-%d-DB", id)})
	fmt.Println(serviveName)
}

func (s *Services) Runner() {
	numWorkers := runtime.NumCPU() / 2
	fmt.Printf("🚀 Starting %d workers...\n", numWorkers)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go s.runWorker(i, &wg)
	}
	wg.Wait()
	fmt.Println("🏁 All workers completed!\n")
}

func main() {
	s := InitServices()
	s.Runner()
	fmt.Println("📦 Final Services List:")
	for name, srv := range s.ServiceList {
		fmt.Printf("📦 %s: %s\n", name, srv.Name)
	}

	srv, ok := s.GetOneService("Service-1")
	if ok {
		fmt.Println("\n🔍 Found single service:", srv.Name)
	} else {
		fmt.Println("\n❌ Service not found")
	}
}
