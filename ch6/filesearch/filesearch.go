package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go fileSearch(os.Args[1], os.Args[2], &wg)
	wg.Wait()
}

func fileSearch(dir, filename string, wg *sync.WaitGroup) {
	// Reads all files from the directory given to the function
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		// Joins each file to the directory. e.g 'cat.jpg' becomes /home/pics/cat.jpg
		fpath := filepath.Join(dir, file.Name())
		if strings.Contains(file.Name(), filename) {
			fmt.Println(fpath)
		}
		// If it's a dir, adds 1 to the waitgroup before starting a new goroutine
		if file.IsDir() {
			wg.Add(1)
			// Creates a goroutine recursively, searching in the new dir
			go fileSearch(fpath, filename, wg)
		}
	}
	// Marks Done() on the waitgroup after processing all files.
	wg.Done()
}
