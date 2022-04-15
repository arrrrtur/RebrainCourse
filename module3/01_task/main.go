//package main
package _1_task

import (
	"fmt"
	"github.com/huandu/xstrings"
	"module3/01_task/wordz"
)

func main() {
	city, digit := wordz.Random()
	fmt.Println(city, digit)

	fmt.Println(xstrings.Shuffle(city))
}
