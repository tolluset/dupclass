package main

func main() {
    searchDir := getSearchDirectory()
    
    printHeader(searchDir)
    duplicates := findDuplicateClasses(searchDir)
    handleResults(duplicates)
}