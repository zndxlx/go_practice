package main


import "fmt"

type object func(int32) int32

func logDecorator(obj object)object{
    return func(input int32)int32{
        fmt.Printf("input=%d\n",input)
        r := obj(input)
        fmt.Printf("output=%d\n", r)
        return r
    }
}

func doubule(n int32) int32 {
    return n * 2
}

func main() {
    f := logDecorator(doubule)
    f(4)
}