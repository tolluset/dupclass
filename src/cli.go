package main

import "os"

// Get search directory from command line args or use default
func getSearchDirectory() string {
    if len(os.Args) > 1 {
        return os.Args[1]
    }
    return "."
}