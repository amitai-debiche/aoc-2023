package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
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
    nums := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

    for scanner.Scan() {
        line := scanner.Text()
        start, end := 0, 0
        startIndex, endIndex := 999, -999
        in_line :=  [9]int{-999,-999,-999,-999,-999,-999,-999,-999,-999}
        for i:=0; i < len(nums); i++ {
            in_line[i] = strings.Index(line, nums[i])
        }

        end_line :=  [9]int{-999,-999,-999,-999,-999,-999,-999,-999,-999}
         for i:=0; i < len(nums); i++ {
            end_line[i] = strings.LastIndex(line, nums[i])
        }


//       fmt.Println(strings.Index(line, nums[0]))

//      fmt.Println(in_line)
        minWordIndex := 999
        for _, v := range in_line {
            if v < minWordIndex && v != -999 && v != -1{
                minWordIndex = v
            }
        }

        maxWordIndex := -111
        for _, v := range end_line {
            if v > maxWordIndex && v != -1 {
                maxWordIndex = v
            }
        }

        minWordVal := 0
        maxWordVal := 0
        for index, num  := range in_line{
            if num == minWordIndex {
                minWordVal = index + 1
            }        
        }

        for index, num := range end_line {
            if num  == maxWordIndex {
                maxWordVal = index + 1
            }
        }

 //      fmt.Println(minWordVal, maxWordVal)
        runeArr := []rune(line) 
        for j := 0; j < len(runeArr) ; j++ {
            if int(runeArr[j]-'0') >= 0 && int(runeArr[j]-'0') <= 9 {
                if i := int(runeArr[j]-'0') ; start == 0 {
                    start = i 
                    end = i
                    startIndex, endIndex = j, j
                } else {
                    end = i 
                    endIndex = j
                }
            }
        }

        if maxWordVal == 0 {
            maxWordVal = minWordVal
        }
        if maxWordIndex > endIndex {
            end = maxWordVal
        }
        if minWordIndex < startIndex {
            start = minWordVal
        }
//        fmt.Println(start, end)
        calibration_values += (start*10)+end
//        fmt.Println(calibration_values)
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
    
    return calibration_values;
}
