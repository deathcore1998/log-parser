package parser

import "strings"

type LogEntry struct {
	Date    string
	Level   string
	Client  string
	Message string
}

func ParseLine(line string) LogEntry {

	logEntry := LogEntry{}

	// find data
	dataEnd := strings.Index(line, "]")
	if dataEnd == -1 {
		return LogEntry{}
	}

	logEntry.Date = line[1:dataEnd]
	line = strings.TrimSpace(line[dataEnd+1:])

	if len(line) < 2 || line[0] != '[' {
		return logEntry
	}

	// find level
	levelEnd := strings.Index(line, "]")
	if levelEnd == -1 {
		return logEntry
	}

	logEntry.Level = line[1:levelEnd]
	line = strings.TrimSpace(line[levelEnd+1:])

	// find client
	clientPrefix := "[client "
	if strings.HasPrefix(line, clientPrefix) {
		endClient := strings.Index(line, "]")
		if endClient != -1 {
			logEntry.Client = line[len(clientPrefix):endClient]
			line = strings.TrimSpace(line[endClient+1:])
		}
	}
	logEntry.Message = line
	return logEntry
}
