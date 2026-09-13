package parser

type LogEntry struct {
	Data    string
	Level   string
	Client  string
	Message string
}

func ParseLine(line string) LogEntry {
	return LogEntry{}
}
