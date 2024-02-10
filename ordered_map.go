package orderedmap

type OrderedMap[K comparable, V any] []Pair[K, V]

type Pair[K comparable, V any] struct {
	Key K
	Val V
}

func New[K comparable, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{}
}

func (m OrderedMap[K, V]) Has(key K) bool {
	_, ok := m.Get(key)
	return ok
}

func (m OrderedMap[K, V]) Get(key K) (V, bool) {
	for _, p := range m {
		if p.Key == key {
			return p.Val, true
		}
	}
	var zero V
	return zero, false
}

func (m *OrderedMap[K, V]) Set(key K, val V) {
	if m == nil {
		*m = OrderedMap[K, V]{}
	}

	for i, p := range *m {
		if p.Key == key {
			(*m)[i].Val = val
			return
		}
	}
	*m = append(*m, Pair[K, V]{key, val})
}

func (m *OrderedMap[K, V]) Delete(key K) {
	if m == nil || len(*m) == 0 {
		return
	}

	for i, p := range *m {
		if p.Key == key {
			*m = append((*m)[:i], (*m)[i+1:]...)
			return
		}
	}
}

func (m OrderedMap[K, V]) Keys() []K {
	if m == nil {
		return nil
	}
	keys := make([]K, len(m))
	for i, p := range m {
		keys[i] = p.Key
	}
	return keys
}

func (m OrderedMap[K, V]) Values() []V {
	if m == nil {
		return nil
	}
	values := make([]V, len(m))
	for i, p := range m {
		values[i] = p.Val
	}
	return values
}

func (m *OrderedMap[K, V]) Clear() {
	if m == nil || len(*m) == 0 {
		return
	}

	*m = []Pair[K, V]{}
}

func (m OrderedMap[K, V]) Copy() OrderedMap[K, V] {
	if m == nil {
		return nil
	}

	return append([]Pair[K, V]{}, m...)
}
