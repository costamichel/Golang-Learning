package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Number  *int   `json:"n,omitempty"` // if pointer is null, will be ignored at Marshal (pointer = nil is the "zero" value to pointers)
	Name    string `json:"name"`
	IsAlive bool   `json:"is_alive,omitempty"` // omitempty ignores field in JSON is value is zero(start value)
}

func main() {
	var n001 *int
	n001 = new(int)
	*n001 = 1
	var n456 *int
	players := []Player{
		{Number: n456, Name: "Gi-Hun", IsAlive: true},
		{Number: n001, Name: "Leader", IsAlive: true},
		{Name: "Dae Ho", IsAlive: false},
	}

	jsonData, _ := json.Marshal(players)
	fmt.Println(string(jsonData))
}
