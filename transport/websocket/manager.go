package websocket

import (
	"sync"
)

// Manager
type Manager struct {
	mu sync.RWMutex

	connections map[uint64]Connection
}

// NewManager
func NewManager() *Manager {
	return &Manager{
		connections: make(map[uint64]Connection),
	}
}

// Add
func (m *Manager) Add(conn Connection) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.connections[conn.GetID()] = conn

}

// Del
func (m *Manager) Remove(conn Connection) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.connections, conn.GetID())
}

// Get
func (m *Manager) Get(cid uint64) (Connection, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if conn, ok := m.connections[cid]; ok {
		return conn, nil
	}
	return nil, ErrConnNotFound
}

// Len
func (m *Manager) Len() int {
	return len(m.connections)
}

// Clear
func (m *Manager) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, conn := range m.connections {
		conn.Stop()
	}
}

// Range
func (m *Manager) Range(f ConnHandlerFunc) (err error) {
	for cid, conn :=  range m.connections {
		err = f(cid, conn)
	}

	return err
}