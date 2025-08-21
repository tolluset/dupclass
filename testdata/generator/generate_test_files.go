package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Common CSS class names for realistic testing
var commonClasses = []string{
	"container", "wrapper", "header", "footer", "nav", "sidebar", "content", "main",
	"button", "btn", "link", "text", "title", "subtitle", "description", "card",
	"modal", "popup", "dropdown", "menu", "list", "item", "row", "col", "grid",
	"flex", "center", "left", "right", "top", "bottom", "hidden", "visible",
	"active", "inactive", "disabled", "enabled", "primary", "secondary", "danger",
	"success", "warning", "info", "loading", "error", "placeholder", "overlay",
}

// Generate duplicate class names for testing
func generateDuplicateClasses() []string {
	var classes []string
	
	// Generate some base classes
	for i := 0; i < 50; i++ {
		base := commonClasses[rand.Intn(len(commonClasses))]
		classes = append(classes, fmt.Sprintf("%s-%d", base, rand.Intn(10)))
	}
	
	// Add some intentional duplicates
	duplicates := []string{
		"duplicate-class-1", "duplicate-class-2", "duplicate-class-3",
		"test-duplicate", "shared-component", "common-style",
	}
	
	classes = append(classes, duplicates...)
	return classes
}

// Generate React component content with specified size
func generateComponentContent(filename string, targetSizeKB int) string {
	classes := generateDuplicateClasses()
	componentName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	componentName = strings.Title(strings.ReplaceAll(componentName, "_", ""))
	
	var content strings.Builder
	
	// Basic component structure
	content.WriteString(fmt.Sprintf(`import React from 'react';

interface %sProps {
  className?: string;
  children?: React.ReactNode;
}

const %s: React.FC<%sProps> = ({ className, children }) => {
  return (
    <div className={` + "`${className || ''} ", componentName, componentName))
	
	// Add classes to reach target size
	currentSize := content.Len()
	targetSize := targetSizeKB * 1024
	
	classIndex := 0
	for currentSize < targetSize {
		if classIndex >= len(classes) {
			classIndex = 0
		}
		
		content.WriteString(classes[classIndex] + " ")
		classIndex++
		currentSize = content.Len()
		
		// Add some structure every 1000 characters
		if currentSize%1000 == 0 {
			content.WriteString(`}>
      <div className="`)
		}
	}
	
	// Close the component
	content.WriteString(`"}>
        {children}
        <div className="nested-component">
          <span className="duplicate-class-1">Test content</span>
          <span className="duplicate-class-2">More test content</span>
          <span className="shared-component">Shared content</span>
        </div>
      </div>
    </div>
  );
};

export default ` + componentName + ";\n")
	
	return content.String()
}

// Generate test files
func generateTestFiles(outputDir string, fileCount int, fileSizeKB int) error {
	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}
	
	fmt.Printf("Generating %d test files (%dKB each) in %s...\n", fileCount, fileSizeKB, outputDir)
	
	for i := 0; i < fileCount; i++ {
		var filename string
		if i%2 == 0 {
			filename = filepath.Join(outputDir, fmt.Sprintf("component_%d.tsx", i))
		} else {
			filename = filepath.Join(outputDir, fmt.Sprintf("component_%d.jsx", i))
		}
		
		content := generateComponentContent(filename, fileSizeKB)
		
		if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %v", filename, err)
		}
		
		if (i+1)%100 == 0 {
			fmt.Printf("Generated %d/%d files...\n", i+1, fileCount)
		}
	}
	
	fmt.Printf("Successfully generated %d test files!\n", fileCount)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run generate_test_files.go <command> [options]")
		fmt.Println("Commands:")
		fmt.Println("  small    - Generate 100 files, 1KB each")
		fmt.Println("  medium   - Generate 1000 files, 10KB each")
		fmt.Println("  large    - Generate 5000 files, 50KB each")
		fmt.Println("  huge     - Generate 10000 files, 100KB each")
		fmt.Println("  custom <count> <sizeKB> - Custom file count and size")
		return
	}
	
	rand.Seed(time.Now().UnixNano())
	
	command := os.Args[1]
	var fileCount, fileSizeKB int
	var outputDir string
	
	switch command {
	case "small":
		fileCount, fileSizeKB = 100, 1
		outputDir = "../generated/test_files_small"
	case "medium":
		fileCount, fileSizeKB = 1000, 10
		outputDir = "../generated/test_files_medium"
	case "large":
		fileCount, fileSizeKB = 5000, 50
		outputDir = "../generated/test_files_large"
	case "huge":
		fileCount, fileSizeKB = 10000, 100
		outputDir = "../generated/test_files_huge"
	case "custom":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run generate_test_files.go custom <count> <sizeKB>")
			return
		}
		var err error
		fileCount, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid file count: %v\n", err)
			return
		}
		fileSizeKB, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("Invalid file size: %v\n", err)
			return
		}
		outputDir = fmt.Sprintf("../generated/test_files_custom_%d_%dkb", fileCount, fileSizeKB)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		return
	}
	
	if err := generateTestFiles(outputDir, fileCount, fileSizeKB); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}