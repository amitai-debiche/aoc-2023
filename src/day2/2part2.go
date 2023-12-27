package main

import (
   "bufio"
   "os"
   "fmt"
   "strings"
   "strconv"
)


func main() {

    fmt.Println(grabBag())
}


func grabBag() int {
    total := 0

    file, err := os.Open("input1.txt")
    if err != nil {
        fmt.Println("Fatal Error")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()[5:]
        colon_split := strings.Split(line, ":")
        
        bag_subset := strings.Split(colon_split[1], ";")

        green_max := 0
        blue_max := 0  
        red_max := 0  
        for _, v := range bag_subset {
        
            picks := strings.Split(v, ",")

            for _, v := range picks {
                elf_pick := strings.Split(v, " ")
                number_taken := stringtoI(elf_pick[1])
                color := elf_pick[2]

                if color == "green" && number_taken > green_max {
                    green_max = number_taken
                }
                if color == "blue" && number_taken > blue_max {
                    blue_max  = number_taken
                }
                if color == "red" && number_taken > red_max  {
                    red_max  = number_taken
                }
            }
        }
        color_power := red_max * green_max * blue_max
        
        total += color_power 
//        fmt.Println(bag_picks)


    }
    return total 
}

func stringtoI(id string) int {
    i, err := strconv.Atoi(id)
    if err != nil {
        panic(err)
    }

    return i
} 
