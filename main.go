package main

import (
	"fmt"

	"github.com/hash_go/hashmap"
)

func main() {
	hm := hashmap.NewHashMap(16)
	hm.Set("hello", "world")
	value := map[string]string{"a": "b"}
	hm.Set(1, value)

	x := hm.Get(1)

	if m, ok := x.(map[string]string); ok {
		if v, exists := m["a"]; exists {
			fmt.Println(v)
		}
	}

}
