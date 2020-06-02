package main

import "fmt"

type Object struct {
    name string
    age  int32
}

type Pool chan *Object

func New(total int) Pool {
    p := make(Pool, total)
    
    for i := 0; i < total; i++ {
        p <- &Object{name: "laden", age: 18}
    }
    
    return p
}

func main() {
    p := New(2)
    
    select {
    case obj := <-p:
        fmt.Printf("obj=%v\n", obj)
        
        p <- obj // 可以用完后又插入到管道
    default:
        // No more objects left — retry later or fail
        return
    }
    
}
