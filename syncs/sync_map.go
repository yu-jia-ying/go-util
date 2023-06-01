package syncs

import (
	"sync"
)

type SynMap[K any, V any] struct {
	syn sync.Map
}

func (m *SynMap[K, V]) Get(k K) (V, bool) {
	anyValues, ok := m.syn.Load(k)
	if ok {
		return anyValues.(V), ok
	}
	var r V
	return r, ok
}

func (m *SynMap[K, V]) Set(k K, v V) {
	m.syn.Store(k, v)
}

func (m *SynMap[K, V]) Remove(k K) {
	m.syn.Delete(k)
}

func (m *SynMap[K, V]) Range(f func(k K, v V) bool) {
	m.syn.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}
