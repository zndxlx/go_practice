package main

import (
    "strings"
    "fmt"
)

func test_builder() {
    b := strings.Builder{}
    b.WriteString("hello")
    b.WriteString(",world")
    fmt.Println(b.String())
}

func main() {
    test_builder()
}
