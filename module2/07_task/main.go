package _7_task

//package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("07_task/data/in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	outputFile, err := os.OpenFile("07_task/data/out.txt", os.O_APPEND|os.O_CREATE, 777)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		i++
		line := strings.Split(scanner.Text(), "|")
		Name, Address, City := line[0], line[1], line[2]
		if Name == "" || Address == "" || City == "" {
			panic(fmt.Sprintf("parse error: empty field on string %d", i))
		}
		_, err := outputFile.WriteString(
			fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", i, Name, Address, City),
		)
		if err != nil {
			panic(err)
		}

	}
}
