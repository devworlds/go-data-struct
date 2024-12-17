package main

import (
	"sync"
)

type Store struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStore() *Store {
	return &Store{data: map[string]string{}}
}

func (s *Store) Set(key, data string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = data
}

func (s *Store) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	data, ok := s.data[key]
	if !ok {
		return ""
	}
	return data
}
