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
    data, err := os.Open("1_sample.txt")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(data)
    
    for scanner.Scan() {
        line := scanner.Text()
        start, end := 0, 0
        runeArr := []rune(line) 
        for j := 0; j < len(runeArr) ; j++ {
            if int(runeArr[j]-'0') >= 0 && int(runeArr[j]-'0') <= 9 {
                if i := int(runeArr[j]-'0') ; start == 0 {
                    start = i 
                    end = i
                } else {
                    end = i
                }
            }
        }
        calibration_values += (start*10)+end
        fmt.Println(calibration_values)
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
    
    return calibration_values;
}
