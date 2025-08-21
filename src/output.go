package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

const (
    colorRed    = "\033[0;31m"
    colorGreen  = "\033[0;32m"
    colorYellow = "\033[1;33m"
    colorReset  = "\033[0m"
)

// Print application header
func printHeader(searchDir string) {
    fmt.Printf("🔍 Checking for duplicate Tailwind classes in: %s\n", searchDir)
    fmt.Println("================================================")
}

// Print success message
func printSuccess() {
    fmt.Println("================================================")
    fmt.Printf("%s✓ No duplicate classes found!%s\n", colorGreen, colorReset)
}

// Handle results and exit appropriately  
func handleResults(duplicates []Duplicate) {
    if len(duplicates) > 0 {
        printResults(duplicates)
        os.Exit(1)
    } else {
        printSuccess()
    }
}

// Print duplicate class results
func printResults(duplicates []Duplicate) {
    for _, d := range duplicates {
        relPath, _ := filepath.Rel(".", d.File)
        fmt.Printf("%s✗%s %s:%d\n", colorRed, colorReset, relPath, d.Line)
        fmt.Printf("  %sDuplicate classes:%s %s\n", colorYellow, colorReset, strings.Join(d.Duplicates, " "))
        
        displayClasses := d.Classes
        if len(displayClasses) > 100 {
            displayClasses = displayClasses[:100] + "..."
        }
        fmt.Printf("  %sFull className:%s %s\n", colorYellow, colorReset, displayClasses)
        fmt.Println()
    }
    fmt.Println("================================================")
    fmt.Printf("%s✗ Found %d line(s) with duplicate classes%s\n", colorRed, len(duplicates), colorReset)
}