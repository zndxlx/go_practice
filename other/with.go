package main

import (
    "fmt"
)
//Option模式，在有字段需要指定默认值,时候有用, 在参数比较多，保证设置参数没有先后顺序时候有用
//扩展设置时候也比较方便
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
    p = &Person{}  //可以填写默认值
    
    for _, option := range options {
        option(p)
    }
    
    return p
}

func main()  {
    p := NewPerson(WithName("lixin"), WithAge(18))
    
    fmt.Printf("p=%+v", p)
}