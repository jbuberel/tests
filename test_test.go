package main

import (
	"bufio"
	"os"
  "log"
  "math/rand"

	"testing"
)


func BenchmarkStr(b *testing.B) {
  words, _ := readLines("wordlist.txt")
  rand.Seed(37)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
    left := words[rand.Intn(len(words))]
    right := words[rand.Intn(len(words))]
    log.Printf("[%v] [%v]\n", left, right)
    substring(left, right)

	}
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
