package main

import (
    "fmt"

    go_say_hello "github.com/DhandyPF/go-say-hello"
)

func main() {
        // variable "result" buat nampung return dari SayHello()
    result := go_say_hello.SayHello()
    fmt.Println(result)
}