[![Go Reference](https://pkg.go.dev/badge/github.com/nicolasparada/go-omap.svg)](https://pkg.go.dev/github.com/nicolasparada/go-omap)

# Golang Ordered Map

Golang Ordered Map is a `map` data structure that maintains the order of the keys.
It also supports JSON and YAML marshalling.

## Installation

```bash
go get github.com/nicolasparada/go-omap
```

## Usage

```go
package main

import (
	omap "github.com/nicolasparada/go-omap"
)

func main() {
	data := []byte(`{ "name": "John", "age": 30, "active": true }`)

	var unordered map[string]any{}
	if err := json.Unmarshal(data, &unordered); err != nil {
		panic(err)
	}

	var ordered omap.Map[string, any]
	if err := json.Unmarshal(data, &ordered); err != nil {
		panic(err)
	}

	json.NewEncoder(os.Stdout).Encode(unordered) // will print in undefined order
	json.NewEncoder(os.Stdout).Encode(ordered) // will always print: {"name":"John","age":30,"active":true}
}
```
