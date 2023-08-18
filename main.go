package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func SearchandReplace(filename, search, replace string) error {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), search, replace, -1)
		lines = append(lines, line)
	}

	return ReplaceText(filename, lines)
}

func ReplaceText(filename string, lines []string) error {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()
	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		check(err)
	}
	return nil
}

func main() {
	arglen := len(os.Args[1:])
	if arglen < 3 {
		fmt.Println("Filename, Search, and Replace needed..")
		return
	}

	filename := os.Args[1]
	search := os.Args[2]
	replace := os.Args[3]
	err := SearchandReplace(filename, search, replace)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Search and replace completed successfully.")
	}
}
