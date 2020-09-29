package main

import (
        "fmt"
        "strings"
)

func main() {

        var input_str string

        fmt.Printf("Enter a string: ")
        fmt.Scanf("%s", &input_str)

        all_lower := strings.ToLower(input_str)
        cond_i := strings.HasPrefix(all_lower, "i")
        cond_a := strings.Contains(all_lower, "a")
        cond_n := strings.HasSuffix(all_lower, "n")
        found := cond_i && cond_a && cond_n

        if found {
                fmt.Println("Found!")
        } else {
                fmt.Println("Not Found!")
        }

}
