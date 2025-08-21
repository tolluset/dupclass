package main

import (
    "bufio"
    "os"
    "path/filepath"
    "strings"
)

// Load gitignore patterns from .gitignore file
func loadGitignorePatterns(rootDir string) []string {
    gitignorePath := filepath.Join(rootDir, ".gitignore")
    file, err := os.Open(gitignorePath)
    if err != nil {
        return nil // No .gitignore file
    }
    defer file.Close()

    var patterns []string
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        // Skip empty lines and comments
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }
        patterns = append(patterns, line)
    }
    
    return patterns
}

// Check if path should be ignored based on gitignore patterns
func shouldIgnore(path string, patterns []string) bool {
    for _, pattern := range patterns {
        if matchesPattern(path, pattern) {
            return true
        }
    }
    return false
}

// Simple pattern matching without external dependencies
func matchesPattern(path, pattern string) bool {
    // Remove leading slash
    if strings.HasPrefix(pattern, "/") {
        pattern = pattern[1:]
    }
    
    // Directory patterns (ending with /)
    if strings.HasSuffix(pattern, "/") {
        dirPattern := pattern[:len(pattern)-1]
        return strings.Contains(path, "/"+dirPattern+"/") || 
               strings.HasSuffix(path, "/"+dirPattern) ||
               strings.HasPrefix(path, dirPattern+"/")
    }
    
    // File patterns
    return strings.Contains(path, pattern) || 
           strings.HasSuffix(path, "/"+pattern) ||
           filepath.Base(path) == pattern
}