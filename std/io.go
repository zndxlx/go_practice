package main

import (
    "strings"
    "io"
    "os"
    "fmt"
    "bytes"
)

func test_copyn(){
    r := strings.NewReader("some word to reader")
    written, err := io.CopyN(os.Stdout, r, 5)
    fmt.Println(written, err)
}

func test_copy() {
    r := strings.NewReader("some word ....")
    
    written, err := io.Copy(os.Stdout, r)
    if err != nil {
        fmt.Println(err)
    }else {
        fmt.Println(written)
    }
}


func test_pipe(){
    r, w := io.Pipe()
    
    go func() {
        fmt.Fprint(w, "sonme txt ....")
        w.Close()
    }()
    
    buf := new(bytes.Buffer)
    buf.ReadFrom(r)
    fmt.Println(buf.String())
}


func test_writeString(){
    io.WriteString(os.Stdout, "hahahhah")
}

func main(){
    //copy()
    test_writeString()
}