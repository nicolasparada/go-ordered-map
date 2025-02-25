package omap

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestMap_Has(t *testing.T) {
	m := Map[string, any]{
		{"name", "John"},
		{"age", 30},
	}
	assert.True(t, m.Has("name"))
	assert.True(t, m.Has("age"))
	assert.False(t, m.Has("active"))
}

func TestMap_Get(t *testing.T) {
	m := Map[string, any]{
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

func TestMap_Set(t *testing.T) {
	m := Map[string, any]{}
	m.Set("name", "John")
	m.Set("age", 30)

	assert.Equal(t, Map[string, any]{
		{"name", "John"},
		{"age", 30},
	}, m)
}

func TestMap_Delete(t *testing.T) {
	m := Map[string, any]{
		{"name", "John"},
		{"age", 30},
	}

	m.Delete("name")

	assert.Equal(t, Map[string, any]{
		{"age", 30},
	}, m)
}
