package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
    
    defer cancel()
    
    select {
    case <-time.After(2 * time.Second):
        fmt.Println("oversleep")
    case <-ctx.Done():
        fmt.Println(ctx.Err())
    }
    
}
