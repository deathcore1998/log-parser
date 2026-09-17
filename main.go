package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/deathcore1998/log-parser/parser"
	"github.com/deathcore1998/log-parser/stats"
)

func main() {
	filePath := flag.String("file", "data/Apache_2k.log", "Path to the log file")
	levelFilter := flag.String("level", "", "Filter by level (error, notice)")

	flag.Parse()

	logFile, err := os.Open(*filePath)
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
		if log.Level == "" {
			continue
		}

		if *levelFilter != "" && *levelFilter != log.Level {
			continue
		}

		levelCount[log.Level]++
		messageCount[log.Message]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading")
		os.Exit(1)
	}

	stats.PrintStats(levelCount, messageCount)
}
