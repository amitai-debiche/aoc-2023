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
    red := 12
    green := 13
    blue := 14
    total := 0

    file, err := os.Open("input1.txt")
    if err != nil {
        fmt.Println("Fatal Error")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        ignore_game := false
        line := scanner.Text()[5:]
        colon_split := strings.Split(line, ":")
        
        bag_subset := strings.Split(colon_split[1], ";")

        for _, v := range bag_subset {

            picks := strings.Split(v, ",")

            for _, v := range picks {
                elf_pick := strings.Split(v, " ")
                number_taken := stringtoI(elf_pick[1])
                color := elf_pick[2]

                if color == "green" && number_taken > green{
                    ignore_game = true
                }
                if color == "blue" && number_taken > blue{
                    ignore_game = true
                }
                if color == "red" && number_taken > red {
                    ignore_game = true
                }
            }
        }
        
        if !ignore_game {
            game_id := stringtoI(colon_split[0])
            total += game_id
        }
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
