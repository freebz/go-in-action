// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
    "fmt"

    "counters"
)

// main is the entry point for the application.
func main() {
    // Create a variable of the unexported type and initialize
    // the value to 10.
    counter := counters.New(10)

    fmt.Printf("Counter: %d\n", counter)
}
