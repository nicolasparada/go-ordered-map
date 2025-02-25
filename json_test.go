package omap

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
)

func TestMap_MarshalJSON(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Second)
	got, err := json.Marshal(Map[string, any]{
		{"name", "John"},
		{"age", 30},
		{"active", true},
		{"last_access_time", now},
	})
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"John","age":30,"active":true,"last_access_time":"`+now.Format(time.RFC3339Nano)+`"}`, string(got))
}

func TestMap_UnmarshalJSON(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Second)
	nowStr := now.Format(time.RFC3339Nano)
	var got Map[string, any]
	err := json.Unmarshal([]byte(`{"name":"John","age":30,"active":true,"last_access_time":"`+nowStr+`"}`), &got)
	assert.NoError(t, err)

	assert.Equal(t, Map[string, any]{
		{"name", "John"},
		{"age", float64(30)}, // json always unmarshals numbers as float64
		{"active", true},
		{"last_access_time", nowStr}, // json doesn't detect time.Time
	}, got)
}
