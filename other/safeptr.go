package main

import (
    "unsafe"
    "fmt"
)

var x struct {
    a bool
    b int16
    c []int
}

func main() {
    // 和 pb := &x.b 等价
    pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
    *pb = 42
    fmt.Println(x.b) // "42"
}

