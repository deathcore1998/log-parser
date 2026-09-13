package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/deathcore1998/log-parser/parser"
)

const (
	logPath = "data/Apache_2k.log"
	maxLine = 10
)

func main() {
	logFile, err := os.Open(logPath)
	if err != nil {
		fmt.Println("Error open file!", err)
		os.Exit(1)
	}

	defer logFile.Close()

	scanner := bufio.NewScanner(logFile)
	countReadLine := 0

	for scanner.Scan() && countReadLine < maxLine {

		log := parser.ParseLine(scanner.Text())

		fmt.Println(log)

		//fmt.Println(scanner.Text())
		countReadLine++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading")
		os.Exit(1)
	}
}
