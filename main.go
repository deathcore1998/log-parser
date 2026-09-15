package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/deathcore1998/log-parser/parser"
	"github.com/deathcore1998/log-parser/stats"
)

const (
	logPath = "data/Apache_2k.log"
)

func main() {
	logFile, err := os.Open(logPath)
	if err != nil {
		fmt.Println("Error open file!", err)
		os.Exit(1)
	}

	defer logFile.Close()

	levelCount := make(map[string]int)
	messageCount := make(map[string]int)

	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {

		log := parser.ParseLine(scanner.Text())
		levelCount[log.Level]++
		messageCount[log.Message]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading")
		os.Exit(1)
	}

	stats.PrintStats(levelCount, messageCount)
}
