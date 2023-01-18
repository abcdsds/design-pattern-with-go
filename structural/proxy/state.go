package proxy

import "sync"

type retentionState struct {
	m    map[string]bool
	lock *sync.RWMutex
}

func (s *retentionState) exists(name string) bool {
	_, ok := s.m[name]

	return ok
}

func (s *retentionState) delete(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.m, name)
}

func (s *retentionState) add(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m[name] = true
}
