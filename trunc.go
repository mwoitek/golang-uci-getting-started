package main

import "fmt"

func main() {

        var num_int int64
        var num_float float64

        fmt.Printf("Enter a floating point number: ")
        fmt.Scanf("%f", &num_float)
        num_int = int64(num_float)
        fmt.Printf("Number after truncation: %d\n", num_int)

}
