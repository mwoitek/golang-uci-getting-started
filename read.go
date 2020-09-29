package main

import (
        "bufio"
        "fmt"
        "os"
        "strings"
)

type name struct {
        fname [20]byte
        lname [20]byte
}

func new_name(fn string, ln string) *name {

        var fn_byte [20]byte
        var ln_byte [20]byte

        copy(fn_byte[:], fn)
        copy(ln_byte[:], ln)

        n := name{fname: fn_byte, lname: ln_byte}
        return &n

}

func main() {

        fmt.Print("\nEnter the name of the text file: ")
        reader := bufio.NewReader(os.Stdin)
        temp_str, _ := reader.ReadString('\n')
        file_name := strings.TrimSuffix(temp_str, "\n")

        file, err1 := os.Open(file_name)
        defer file.Close()

        if err1 == nil {

                var err2 error
                var fn string
                var ln string
                var n *name
                name_slice := make([]name, 0)
                reader = bufio.NewReader(file)

                for {
                        _, err2 = fmt.Fscanf(reader, "%s %s\n", &fn, &ln)
                        if err2 != nil {
                                break
                        }
                        n = new_name(fn, ln)
                        name_slice = append(name_slice, *n)
                }

                for _, ns := range name_slice {
                        fmt.Printf("\nFirst name: %s\n", ns.fname)
                        fmt.Printf("Last name: %s\n", ns.lname)
                }
                fmt.Println("")

        } else {

                fmt.Println("Could not open the text file!")

        }

}
