package main

import (
    "testing"
)

func TestIsReactFile(t *testing.T) {
    tests := []struct {
        name     string
        filename string
        expected bool
    }{
        {"TSX file", "component.tsx", true},
        {"JSX file", "component.jsx", true},
        {"TypeScript file", "utils.ts", false},
        {"JavaScript file", "script.js", false},
        {"CSS file", "styles.css", false},
        {"Short name", "a.ts", false},
        {"No extension", "README", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := isReactFile(tt.filename)
            if result != tt.expected {
                t.Errorf("isReactFile(%q) = %v, expected %v", tt.filename, result, tt.expected)
            }
        })
    }
}