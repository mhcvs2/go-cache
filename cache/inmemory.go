package cache

import "sync"

type inMemoryCache struct {
	c     map[string][]byte
	mutex sync.RWMutex
	Stat
}

func (s *inMemoryCache) Set(k string, v []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if ov, exist := s.c[k]; exist {
		s.del(k, ov)
	}
	s.c[k] = v
	s.add(k, v)
	return nil
}

func (s *inMemoryCache) Get(k string) ([]byte, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.c[k], nil
}

func (s *inMemoryCache) Del(k string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if v, exist := s.c[k]; exist {
		delete(s.c, k)
		s.del(k, v)
	}
	return nil
}

func (s *inMemoryCache) GetStat() Stat {
	return s.Stat
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{make(map[string][]byte), sync.RWMutex{}, Stat{}}
}
