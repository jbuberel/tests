package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"

	"testing"
)

func BenchmarkStr(b *testing.B) {
	words, _ := readLines("wordlist.txt")
	substrings := 0
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if strings.Contains(words[rand.Intn(len(words))], words[rand.Intn(len(words))]) {
			substrings++
		}
	}
	log.Printf("Total substring matches: %v Total comparisons: %v Percent: %0.5f\n", substrings, b.N, 100.0*float64(substrings)/float64(b.N))
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
