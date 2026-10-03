package redis

import (
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

// manager redis
var manager *Manager
type Manager struct {
	mtx sync.RWMutex

	clients map[string]*redis.Client
}


func (m *Manager)NewClient(name string, c *Config) {
	if name == "" {
		name = "default"
	}

	m.mtx.RLock()
	if _, ok := m.clients[name]; ok {
		m.mtx.Unlock()
		return
	}

	m.mtx.RUnlock()

	m.mtx.Lock()

	defer m.mtx.Unlock()

	rdb, err := NewClient(c)
	if err != nil {
		log.Fatalf("init redis client err:%v", err)
	}

	m.clients[name] = rdb
}

// GetClient

func (m *Manager) GetClient(name string) *redis.Client {
	m.mtx.RLock()
	defer m.mtx.Unlock()

	if c, ok := m.clients[name]; ok {
		return c
	}

	return nil
}


// NewManager

func NewManager() *Manager {
	m := &Manager{
		clients: make(map[string]*redis.Client),
	}
	manager = m
	return m
}

// GetManager

func GetManager() *Manager {
	return manager
}

func GetClient(name string) *redis.Client {
	return manager.GetClient(name)
}

