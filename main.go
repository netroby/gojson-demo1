package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"name":"john", "age":20}`)
	var f map[string]interface{}
	json.Unmarshal(b, &f)
	fmt.Printf("%s", f["name"])
}
