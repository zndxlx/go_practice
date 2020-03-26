package main

import (
    "context"
    "log"
    "net"
    "time"
)

func main(){
    var d net.Dialer
    
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()
    
    conn, err := d.DialContext(ctx, "tcp", "localhost:20000")
    defer conn.Close()
    
    _, err = conn.Write([]byte("hhhhhhhhh"))
    if err != nil {
        log.Fatal(err)
    }
    
    
}
