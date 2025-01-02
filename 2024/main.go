package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	TESTING = false
)

func loadEnv() {
	envBytes, err := os.ReadFile(".env")
	if err != nil {
		log.Fatalf("could not open env file, %v", err)
	}
	envFile := bytes.NewReader(envBytes)
	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		enVar := strings.Split(scanner.Text(), "=")
		os.Setenv(enVar[0], enVar[1])
	}
}

func GetInputFile(day int) {
	loadEnv()

	_, err := os.Stat(fmt.Sprintf("day%d.txt", day))
	if err == nil {
		fmt.Println("File Exists")
		return
	}
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("Cookie", fmt.Sprintf("ru=%s; session=%s", os.Getenv("RU"), os.Getenv("SESSION")))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	file, err := os.OpenFile(fmt.Sprintf("day%d.txt", day), os.O_WRONLY|os.O_CREATE, 0644)
	_, err = fmt.Fprint(file, string(body))
	if err != nil {
		fmt.Println("Error occured while trying to write to file,", err)
	}
}

func ParseInputFile(day int, parseFunc func(line string)) {
	var (
		inputBytes []byte
		err        error
	)
	if TESTING {
		inputBytes, err = os.ReadFile(fmt.Sprintf("day%d.test", day))

	} else{
		inputBytes, err = os.ReadFile(fmt.Sprintf("day%d.txt", day))
	}
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	inputFile := bytes.NewReader(inputBytes)
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		parseFunc(scanner.Text())
	}
}

func main() {
	// Day1()
	// Day2()
	// Day3()
	// Day4()
	Day5()
}
