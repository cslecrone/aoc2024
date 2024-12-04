# aoc2024
Solutions for Advent of Code 2024

# Solution Generation
`setup.go` is a small tool for downloading daily input files and generating blank solution files:
```bash
# setup.go expects that the AOC_SESSION_ID environment variable is set
# It should contain the session_id cookie from AOC
$ go run setup.go -day 2
Created directories and input for day 2
```