package main

// Find duplicate classes in a space-separated string
func findDuplicates(classes string) []string {
    if len(classes) == 0 {
        return nil
    }

    var classList []string
    var duplicates []string
    
    // Manual splitting with minimal memory allocation
    start := 0
    for i, r := range classes + " " {
        if r == ' ' {
            if i > start {
                class := classes[start:i]
                if isDuplicate(class, classList) {
                    if !contains(duplicates, class) {
                        duplicates = append(duplicates, class)
                    }
                } else {
                    classList = append(classList, class)
                }
            }
            start = i + 1
        }
    }
    
    return duplicates
}

// Check if class is already in the list
func isDuplicate(class string, classList []string) bool {
    for _, existing := range classList {
        if existing == class {
            return true
        }
    }
    return false
}

// Check if slice contains item
func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}