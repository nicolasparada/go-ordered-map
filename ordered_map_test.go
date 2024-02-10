package orderedmap

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestOrderedMap_Has(t *testing.T) {
	m := OrderedMap[string, any]{
		{"name", "John"},
		{"age", 30},
	}
	assert.True(t, m.Has("name"))
	assert.True(t, m.Has("age"))
	assert.False(t, m.Has("active"))
}

func TestOrderedMap_Get(t *testing.T) {
	m := OrderedMap[string, any]{
		{"name", "John"},
		{"age", 30},
	}
	{
		v, ok := m.Get("name")
		assert.True(t, ok)
		assert.Equal(t, "John", v)
	}
	{
		v, ok := m.Get("age")
		assert.True(t, ok)
		assert.Equal(t, 30, v)
	}
	{
		_, ok := m.Get("active")
		assert.False(t, ok)
	}
}

func TestOrderedMap_Set(t *testing.T) {
	m := OrderedMap[string, any]{}
	m.Set("name", "John")
	m.Set("age", 30)

	assert.Equal(t, OrderedMap[string, any]{
		{"name", "John"},
		{"age", 30},
	}, m)
}

func TestOrderedMap_Delete(t *testing.T) {
	m := OrderedMap[string, any]{
		{"name", "John"},
		{"age", 30},
	}

	m.Delete("name")

	assert.Equal(t, OrderedMap[string, any]{
		{"age", 30},
	}, m)
}

func TestOrderedMap_Keys(t *testing.T) {
	m := OrderedMap[string, any]{
		{"name", "John"},
		{"age", 30},
	}
	assert.Equal(t, []string{"name", "age"}, m.Keys())
}

func TestOrderedMap_Values(t *testing.T) {
	m := OrderedMap[string, any]{
		{"name", "John"},
		{"age", 30},
	}
	assert.Equal(t, []any{"John", 30}, m.Values())
}

func TestOrderedMap_Clear(t *testing.T) {
	m := OrderedMap[string, any]{
		{"name", "John"},
		{"age", 30},
	}
	m.Clear()
	assert.Equal(t, OrderedMap[string, any]{}, m)
}

func TestOrderedMap_Copy(t *testing.T) {
	m := OrderedMap[string, any]{
		{"name", "John"},
		{"age", 30},
	}

	c := m.Copy()
	assert.Equal(t, m, c)

	c.Set("name", "Jane")
	assert.NotEqual(t, m, c)
}
