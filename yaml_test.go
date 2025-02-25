package omap

import (
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
	"gopkg.in/yaml.v3"
)

func TestMap_MarshalYAML(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Second)
	got, err := yaml.Marshal(Map[string, any]{
		{"name", "John"},
		{"age", 30},
		{"active", true},
		{"last_access_time", now},
	})
	assert.NoError(t, err)
	assert.Equal(t, "name: John\nage: 30\nactive: true\nlast_access_time: "+now.Format(time.RFC3339Nano)+"\n", string(got))
}

func TestMap_UnmarshalYAML(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Second)
	var got Map[string, any]
	err := yaml.Unmarshal([]byte("name: John\nage: 30\nactive: true\nlast_access_time: "+now.Format(time.RFC3339Nano)+"\n"), &got)
	assert.NoError(t, err)

	assert.Equal(t, Map[string, any]{
		{"name", "John"},
		{"age", 30},
		{"active", true},
		{"last_access_time", now},
	}, got)
}
