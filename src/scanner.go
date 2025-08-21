package main

import (
    "os"
    "path/filepath"
    "sync"
)

// Recursively scan directory for React files
func scanDirectory(dir string, fileChan chan<- string) {
    entries, err := os.ReadDir(dir)
    if err != nil {
        return
    }

    for _, entry := range entries {
        path := filepath.Join(dir, entry.Name())
        if entry.IsDir() {
            scanDirectory(path, fileChan)
        } else {
            if isReactFile(entry.Name()) {
                fileChan <- path
            }
        }
    }
}

// Check if file is a React component file
func isReactFile(name string) bool {
    return (len(name) > 4 && name[len(name)-4:] == ".tsx") ||
           (len(name) > 4 && name[len(name)-4:] == ".jsx")
}

// Worker goroutine to process files
func worker(wg *sync.WaitGroup, fileChan <-chan string, duplicateChan chan<- Duplicate) {
    defer wg.Done()
    
    // Reusable buffer for performance
    buf := make([]byte, 0, 1024)
    
    for path := range fileChan {
        checkFile(path, duplicateChan, &buf)
    }
}