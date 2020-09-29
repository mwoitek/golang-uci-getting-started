package main

import (
        "bufio"
        "encoding/json"
        "fmt"
        "os"
        "strings"
)

func main() {

        var personal_info map[string]string
        personal_info = make(map[string]string)
        reader := bufio.NewReader(os.Stdin)

        fmt.Print("Enter a name: ")
        input_str, _ := reader.ReadString('\n')
        personal_info["name"] = strings.TrimSuffix(input_str, "\n")

        fmt.Print("Enter an address: ")
        input_str, _ = reader.ReadString('\n')
        personal_info["address"] = strings.TrimSuffix(input_str, "\n")

        json_obj, _ := json.Marshal(personal_info)
        os.Stdout.Write(json_obj)
        fmt.Println("")

}
