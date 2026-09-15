package parser

import "strings"

type LogEntry struct {
	Data    string
	Level   string
	Client  string
	Message string
}

func ParseLine(line string) LogEntry {

	logEntry := LogEntry{}

	// find data
	dataEnd := strings.Index(line, "]") + 1
	if dataEnd == 0 {
		return LogEntry{}
	}

	logEntry.Data = line[:dataEnd]
	line = strings.TrimSpace(line[dataEnd:])

	if len(line) < 2 || line[0] != '[' {
		return logEntry
	}

	// find level
	levelEnd := strings.Index(line, "]") + 1
	if levelEnd == 0 {
		return logEntry
	}

	logEntry.Level = line[:levelEnd]
	line = strings.TrimSpace(line[levelEnd:])

	// find client
	if strings.HasPrefix(line, "[client ") {
		endClient := strings.Index(line, "]") + 1
		if endClient != 0 {
			logEntry.Client = line[:endClient]
			line = line[endClient:]
		}
	}
	logEntry.Message = line
	return logEntry
}
