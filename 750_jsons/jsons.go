package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}

func main() {
	u := &User{
		Name: "Jacob",
		Age:  24,
	}

	bytes, _ := json.Marshal(u)
	fmt.Println(string(bytes))

	u2 := &User{}
	json.Unmarshal([]byte(`{"name":"Jacob"}`), u2)
	fmt.Println(u2)
}
