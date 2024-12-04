package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const templateSource = "./templates/main.go"

func parseDay() int {
	help := flag.Bool("help", false, "Display help information")
	day := flag.Int("day", 0, "The day to setup")

	flag.Parse()

	if *help {
		fmt.Println("Setup problems for AOC")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *day < 1 || *day > 25 {
		log.Fatalf("Day %d is not valid day in Advent", *day)
	}

	return *day
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	if info.IsDir() {
		return true
	} else {
		return false
	}
}

func createDirectories(day int) {
	basePath := "./day-" + strconv.Itoa(day)
	if dirExists(basePath) {
		log.Fatalf("A directory already exists for day %d\n", day)
	}

	src, err := os.Open(templateSource)

	if err != nil {
		log.Fatal(err)
	}

	defer src.Close()

	for i := 1; i < 3; i++ {
		p := basePath + "/part-" + strconv.Itoa(i)
		fp := p + "/main.go"
		os.MkdirAll(p, 0755)
		dest, err := os.Create(fp)
		if err != nil {
			log.Fatal(err)
		}
		defer dest.Close()

		_, err = io.Copy(dest, src)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func getInput(day int, sessionId string) []byte {
	url := "https://adventofcode.com/2024/day/" + strconv.Itoa(day) + "/input"

	req, _ := http.NewRequest("GET", url, nil)
	cookie := &http.Cookie{
		Name:  "session",
		Value: sessionId,
	}
	req.AddCookie(cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Unable to make a request: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Recieved error status code %d attempting to GET %s", resp.StatusCode, url)
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to parse response body: %s", err)
	}

	return b
}

func writeInput(input []byte, day int) {
	p := "./day-" + strconv.Itoa(day) + "/input.txt"
	f, err := os.Create(p)
	if err != nil {
		log.Fatalf("Unable to create file %s: %s", p, err)
	}
	defer f.Close()

	err = os.WriteFile(p, input, 0755)
	if err != nil {
		log.Fatalf("Unable to write input to %s: %s", p, err)
	}
}

func getSessionId() string {
	sessionId, exists := os.LookupEnv("AOC_SESSION_ID")
	if !exists {
		log.Fatalf("AOC_SESSION_ID environment variable is required.")
	}
	return sessionId
}

func main() {
	sessionId := getSessionId()
	day := parseDay()
	inputBody := getInput(day, sessionId)
	createDirectories(day)
	writeInput(inputBody, day)
	fmt.Printf("Created directories and input for day %d\n", day)
}
