package main

import (
    "bufio"
    "os"
    "strings"
)

// Check file for duplicate className values
func checkFile(path string, duplicateChan chan<- Duplicate, buf *[]byte) {
    file, err := os.Open(path)
    if err != nil {
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // Buffer size optimization
    *buf = (*buf)[:0]
    if cap(*buf) < 64*1024 {
        *buf = make([]byte, 64*1024)
    }
    scanner.Buffer(*buf, 64*1024)
    
    lineNum := 0

    for scanner.Scan() {
        lineNum++
        line := scanner.Text()

        if !strings.Contains(line, "className=") {
            continue
        }

        classes := extractClassName(line)
        if classes == "" {
            continue
        }

        duplicateClasses := findDuplicates(classes)
        if len(duplicateClasses) > 0 {
            duplicateChan <- Duplicate{
                File:       path,
                Line:       lineNum,
                Classes:    classes,
                Duplicates: duplicateClasses,
            }
        }
    }
}

// Extract className value without regex
func extractClassName(line string) string {
    idx := strings.Index(line, "className=")
    if idx == -1 {
        return ""
    }
    
    startIdx := idx + 10 // len("className=")
    if startIdx >= len(line) {
        return ""
    }
    
    nextChar := line[startIdx]
    
    switch nextChar {
    case '"':
        return extractQuoted(line, startIdx+1, '"')
    case '\'':
        return extractQuoted(line, startIdx+1, '\'')
    case '{':
        return extractFromBraces(line, startIdx+1)
    }
    
    return ""
}

// Extract quoted string value
func extractQuoted(line string, start int, quote byte) string {
    if end := strings.IndexByte(line[start:], quote); end != -1 {
        return line[start : start+end]
    }
    return ""
}

// Extract value from braces (template literals or function calls)
func extractFromBraces(line string, start int) string {
    if start >= len(line) {
        return ""
    }
    
    // Check for backtick template literals
    if line[start] == '`' {
        return extractQuoted(line, start+1, '`')
    }
    
    // Check for function call patterns like cn("...") or clsx('...')
    for i := start; i < len(line)-2; i++ {
        if line[i] == '(' {
            j := i + 1
            // Skip whitespace
            for j < len(line) && line[j] == ' ' {
                j++
            }
            if j >= len(line) {
                break
            }
            
            switch line[j] {
            case '"':
                return extractQuoted(line, j+1, '"')
            case '\'':
                return extractQuoted(line, j+1, '\'')
            case '`':
                return extractQuoted(line, j+1, '`')
            }
            break
        }
    }
    
    return ""
}