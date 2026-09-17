package parser

import "testing"

type TestArgs struct {
	name     string
	line     string
	expected LogEntry
}

func TestParseLine(t *testing.T) {

	tests := []TestArgs{
		{
			name: "Notice without client",
			line: "[Sun Dec 04 05:04:03 2005] [notice] jk2_init() Found child 8743 in scoreboard slot 7",
			expected: LogEntry{
				Date:    "Sun Dec 04 05:04:03 2005",
				Level:   "notice",
				Client:  "",
				Message: "jk2_init() Found child 8743 in scoreboard slot 7",
			},
		},
		{
			name: "Error with client",
			line: "[Mon Dec 05 10:26:39 2005] [error] [client 141.153.150.164] Directory index forbidden by rule: /var/www/html/",
			expected: LogEntry{
				Date:    "Mon Dec 05 10:26:39 2005",
				Level:   "error",
				Client:  "141.153.150.164",
				Message: "Directory index forbidden by rule: /var/www/html/",
			},
		},
		{
			name:     "Empty string",
			line:     "",
			expected: LogEntry{},
		},
		{
			name:     "Invalid string",
			line:     "not a log line at all",
			expected: LogEntry{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ParseLine(test.line)
			if result != test.expected {
				t.Errorf("ParseLine(%q)\ngot:  %+v\nwant: %+v", test.line, result, test.expected)
			}
		})
	}
}
