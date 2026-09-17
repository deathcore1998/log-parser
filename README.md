# Log Parser

A simple CLI tool for parsing Apache log files. Written in Go.

## What it does

- Reads an Apache log file
- Parses each line into: date, level, client IP, message
- Counts how many times each level appears (error, notice, etc.)
- Shows the top 10 most frequent messages

## How to run

### Build

go build -o log-parser

### Run with default file

./log-parser

### Run with flags

./log-parser --file=data/Apache_2k.log --level=error

## Flags
| `--file` | `data/Apache_2k.log` | Path to the log file |
| `--level` | `""` (all) | Filter by level (error, notice) |

## Example output

Statistic:
notice: 1405
error: 595
1. workerEnv.init() ok /etc/httpd/conf/workers2.properties:  569
2. mod_jk child workerEnv in error state 6:  369
3. mod_jk child workerEnv in error state 7:  101
...

## Project structure

log-parser/
├── main.go              # entry point, CLI flags
├── parser/
│   └── apache.go        # parsing logic
│   └── apache_test.go   # tests
├── stats/
│   └── stats.go         # statistics output
└── data/
    └── Apache_2k.log    # sample log file

## Tests

go test ./...

## Tech stack

- go 1.27.1
- Standard library only (no external dependencies)