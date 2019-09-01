package cache

type Stat struct {
	Count     int64 `json:"count"`
	KeySize   int64 `json:"key_size"`
	ValueSize int64 `json:"value_size"`
}

func (s *Stat) add(k string, v []byte) {
	s.Count += 1
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
}

func (s *Stat) del(k string, v []byte) {
	s.Count -= 1
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
}
