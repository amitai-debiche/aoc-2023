package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {
 
    fmt.Println(sum())
}



func sum() int {
    calibration_values := 0 
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
    
    fmt.Println("Hi")    
    return calibration_values;
}
