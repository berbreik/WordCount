package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func main() {
	filenames := []string{
		"file1",
		"file2",
		"file3",
	}

	totalWordCount := ConcurrentWordCount(filenames)
	fmt.Println("Total Word Count:", totalWordCount)
}

func WordCount(filename string, Result chan<- int) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error Reading file ", err.Error())
	}
	Word := strings.Fields(string(content))
	Count := len(Word)
	Result <- Count
}

func ConcurrentWordCount(filename []string) int {
	var wg sync.WaitGroup
	Result := make(chan int)
	for _, file := range filename {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			WordCount(file, Result)
		}(file)
	}
	go func() {
		wg.Wait()
		close(Result)
	}()

	totalwordcount := 0
	for words := range Result {
		totalwordcount += words
	}

	return totalwordcount
}
