package _4_task

//package main

import (
	"fmt"
)

func main() {
	items := make(map[string]map[string][]string, 5)

	items["Bob"] = map[string][]string{
		"Random": {"asda", "sdsad", "asdsadas"},
	}

	for name, v := range items {
		fmt.Println(name, ":")
		for key, value := range v {
			fmt.Println(key, value)
		}
	}
}
