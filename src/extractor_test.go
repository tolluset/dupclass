package main

import (
    "testing"
)

func TestExtractClassName(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {
            name:     "Double quotes",
            input:    `<div className="flex items-center">`,
            expected: "flex items-center",
        },
        {
            name:     "Single quotes", 
            input:    `<div className='px-4 py-2'>`,
            expected: "px-4 py-2",
        },
        {
            name:     "Backticks",
            input:    "<div className={`bg-blue-500 text-white`}>",
            expected: "bg-blue-500 text-white",
        },
        {
            name:     "cn function with double quotes",
            input:    `<div className={cn("rounded-lg shadow-md")}>`,
            expected: "rounded-lg shadow-md",
        },
        {
            name:     "clsx function with single quotes",
            input:    `<div className={clsx('border border-gray-300')}>`,
            expected: "border border-gray-300",
        },
        {
            name:     "Empty string",
            input:    `<div className="">`,
            expected: "",
        },
        {
            name:     "No className",
            input:    `<div class="test">`,
            expected: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := extractClassName(tt.input)
            if result != tt.expected {
                t.Errorf("extractClassName(%q) = %q, expected %q", tt.input, result, tt.expected)
            }
        })
    }
}

