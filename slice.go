package main

import (
        "fmt"
        "strconv"
        "strings"
        "sort"
)

func main() {

        var err error
        var input_str string
        var new_int int
        int_slice := make([]int, 0, 3)

        for {

                fmt.Printf("Enter an integer or X to quit: ")
                fmt.Scanf("%s", &input_str)

                if strings.Compare(strings.ToLower(input_str), "x") == 0 {
                        break
                } else {
                        new_int, err = strconv.Atoi(input_str)
                        if err != nil {
                                fmt.Println("Invalid input!")
                        } else {
                                int_slice = append(int_slice, new_int)
                                sort.Ints(int_slice)
                                fmt.Println(int_slice)
                        }
                }

        }

}
