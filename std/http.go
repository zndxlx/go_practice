package main

import "net/http"

func main() {
    
    http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){}))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){})
    http.ListenAndServe(":8899",nil)
}