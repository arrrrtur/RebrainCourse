package _8_task

//package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("08_task/data/in.txt")
	if err != nil {
		fmt.Println("File don't open.")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("File don't close.")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
	}
	fmt.Println("End of file.")
	fmt.Printf("Total strings: %d", i)
}
