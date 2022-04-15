package _6_task

//package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//defer logTime()

	file, err := os.Open("06_task/data/in.txt")
	if err != nil {
		panic(err)
	}

	fileOutput, err := os.OpenFile("06_task/data/out.txt", os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	i := 0
	countOfBytes := 0

	for scanner.Scan() {
		i++
		writeString, err := fileOutput.WriteString(fmt.Sprintf("%d. %v\n", i, scanner.Text()))
		if err != nil {
			return
		}
		countOfBytes += writeString
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	defer func(fileOutput *os.File) {
		err := fileOutput.Close()

		fmt.Printf("out.txt is closes.\n"+
			"Count of written strings - %d.\n"+
			"Count of written bytes - %d",
			i, countOfBytes)

		if err != nil {
			panic(err)
		}
	}(fileOutput)
}

// TODO logTime
