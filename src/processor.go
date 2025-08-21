package main

import (
    "runtime"
    "sync"
)


// Find duplicate classes in the given directory
func findDuplicateClasses(searchDir string) []Duplicate {
    return processFiles(searchDir)
}

// Process files in directory and return duplicates
func processFiles(searchDir string) []Duplicate {
    // Initialize gitignore patterns
    initGitignore(searchDir)
    
    maxWorkers := getOptimalWorkerCount()

    fileChan := make(chan string, 100)
    duplicateChan := make(chan Duplicate, 100)
    var duplicates []Duplicate

    // Result collection goroutine
    done := make(chan bool)
    go func() {
        for d := range duplicateChan {
            duplicates = append(duplicates, d)
        }
        done <- true
    }()

    // Start worker goroutines
    var wg sync.WaitGroup
    for i := 0; i < maxWorkers; i++ {
        wg.Add(1)
        go worker(&wg, fileChan, duplicateChan)
    }

    // Start file scanning
    go func() {
        defer close(fileChan)
        scanDirectory(searchDir, fileChan)
    }()

    // Wait for all workers to complete
    wg.Wait()
    close(duplicateChan)
    <-done

    return duplicates
}

// Get optimal worker count based on CPU cores
func getOptimalWorkerCount() int {
    maxWorkers := runtime.NumCPU()
    if maxWorkers > 8 {
        maxWorkers = 8 // Prevent excessive goroutine overhead
    }
    return maxWorkers
}