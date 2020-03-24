package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func test_writer() {
    w := bufio.NewWriter(os.Stdout)
    fmt.Fprint(w, "hello,")
    fmt.Fprintf(w, "world")
    w.Flush()
    
}

func test_scanner_bytes(){
    scanner := bufio.NewScanner(strings.NewReader("hellorea\nder"))
    
    for scanner.Scan(){
        fmt.Println(len(scanner.Bytes()))
    }
    
    fmt.Println(scanner.Err())
    
}

func test_scanner_lines(){
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan(){
        fmt.Println(scanner.Text())
    }
    fmt.Println(scanner.Err())
}


func test_scanner_words(){
    scanner := bufio.NewScanner(strings.NewReader("now is the winter \n please be happy"))
    scanner.Split(bufio.ScanWords)
    for scanner.Scan(){
        fmt.Println(scanner.Text())
    }
    fmt.Println(scanner.Err())

}

func main() {
    test_scanner_words()
}
