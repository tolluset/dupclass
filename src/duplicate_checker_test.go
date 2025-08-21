package main

import (
    "testing"
)

func TestFindDuplicates(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected []string
    }{
        {
            name:     "No duplicates",
            input:    "flex items-center justify-between",
            expected: nil,
        },
        {
            name:     "Single duplicate",
            input:    "flex flex items-center",
            expected: []string{"flex"},
        },
        {
            name:     "Multiple duplicates",
            input:    "mt-4 mb-4 mt-4 mb-4 px-2",
            expected: []string{"mt-4", "mb-4"},
        },
        {
            name:     "Empty string",
            input:    "",
            expected: nil,
        },
        {
            name:     "Single class",
            input:    "flex",
            expected: nil,
        },
        {
            name:     "Triple duplicate",
            input:    "bg-red-500 bg-red-500 bg-red-500",
            expected: []string{"bg-red-500"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := findDuplicates(tt.input)
            if !slicesEqual(result, tt.expected) {
                t.Errorf("findDuplicates(%q) = %v, expected %v", tt.input, result, tt.expected)
            }
        })
    }
}

func TestIsDuplicate(t *testing.T) {
    classList := []string{"flex", "items-center", "justify-between"}
    
    tests := []struct {
        name     string
        class    string
        expected bool
    }{
        {"Existing class", "flex", true},
        {"Non-existing class", "bg-red-500", false},
        {"Empty class", "", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := isDuplicate(tt.class, classList)
            if result != tt.expected {
                t.Errorf("isDuplicate(%q, %v) = %v, expected %v", tt.class, classList, result, tt.expected)
            }
        })
    }
}

func TestContains(t *testing.T) {
    slice := []string{"apple", "banana", "cherry"}
    
    tests := []struct {
        name     string
        item     string
        expected bool
    }{
        {"Existing item", "banana", true},
        {"Non-existing item", "grape", false},
        {"Empty item", "", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := contains(slice, tt.item)
            if result != tt.expected {
                t.Errorf("contains(%v, %q) = %v, expected %v", slice, tt.item, result, tt.expected)
            }
        })
    }
}

func slicesEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}