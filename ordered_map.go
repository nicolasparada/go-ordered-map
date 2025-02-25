package omap

import "slices"

type Map[K comparable, V any] []Pair[K, V]

type Pair[K comparable, V any] struct {
	Key K
	Val V
}

func (p Pair[K, V]) Equal(other Pair[K, V]) bool {
	return p.Key == other.Key && any(p.Val) == any(other.Val)
}

func New[K comparable, V any]() Map[K, V] {
	return Map[K, V]{}
}

func (om Map[K, V]) Has(key K) bool {
	return slices.ContainsFunc(om, func(p Pair[K, V]) bool {
		return p.Key == key
	})
}

func (om Map[K, V]) Get(key K) (V, bool) {
	for _, p := range om {
		if p.Key == key {
			return p.Val, true
		}
	}

	var zero V
	return zero, false
}

func (om *Map[K, V]) Set(key K, val V) {
	if om == nil {
		return
	}

	for i, p := range *om {
		if p.Key == key {
			(*om)[i].Val = val
			return
		}
	}

	*om = append(*om, Pair[K, V]{key, val})
}

func (om *Map[K, V]) Delete(key K) {
	if om == nil || len(*om) == 0 {
		return
	}

	for i, p := range *om {
		if p.Key == key {
			*om = append((*om)[:i], (*om)[i+1:]...)
			return
		}
	}
}
