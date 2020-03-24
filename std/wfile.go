package main

import (
    "os"
    "fmt"
)


func main() {
    file, e := os.OpenFile("test.txt", os.O_CREATE , 0755)
    if e != nil {
        fmt.Println(e)
        return
    }
    defer file.Close()
    
    file.WriteString("hhhhhhh")
}
