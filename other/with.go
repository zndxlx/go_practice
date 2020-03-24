package main

import (
    "fmt"
)

type Person struct {
    name string
    age int
}

type option func(p* Person)

func WithName(name string) option {
    return  func(p *Person) {
        p.name = name
    }
}

func WithAge(age int) option {
    return func(p *Person) {
        p.age = age
    }
}

func NewPerson(options ...option)(p *Person) {
    p = &Person{}
    
    for _, option := range options {
        option(p)
    }
    
    return p
}

func main()  {
    p := NewPerson(WithName("lixin"), WithAge(18))
    
    fmt.Printf("p=%+v", p)
}