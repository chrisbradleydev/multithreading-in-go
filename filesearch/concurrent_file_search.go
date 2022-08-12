package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches []string
	waitgroup = sync.WaitGroup{}
	lock = sync.Mutex{}
)

func fileSearch(root, filename string) {
	fmt.Println("searching in", root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			waitgroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}
	waitgroup.Done()
}

func main() {
	homeDir, _ := os.UserHomeDir()
	waitgroup.Add(1)
	go fileSearch(homeDir, "README.md")
	waitgroup.Wait()
	for _, file := range matches {
		fmt.Println("matched", file)
	}
}
