package orm

import (
	"sync"
	
	"gorm.io/gorm"
)


var manager *Manager

type Manager struct {
	mtx		sync.RWMutex

	clients 	map[string]*gorm.DB
}


func (m *Manager) init(name string, c *Config) {
	if name == ""{
		name = "default"
	}

	m.mtx.RLock()
	if _, ok := m.clients[name]; ok {
		m.mtx.RUnlock()
		return
	}

	m.mtx.RUnlock()

	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.clients[name] = NewDB(c)
}

// GetClient 
func (m *Manager) GetClient(name string) *gorm.DB {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	if c, ok := m.clients[name]; ok {
		return c
	}

	return nil
}

// NewManager
func NewManager() *Manager {
	m := &Manager {
		clients: make(map[string]*gorm.DB),
	}
	manager = m
	return m
}

// GetManager
func GetManager() *Manager{
	return manager
}

func Client(name string) *gorm.DB {
	return manager.GetClient(name)
}