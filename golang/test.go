package main

import "fmt"

func main() {
    for i := 0; i < 100 ; i++ {
        if i % 2 == 0 {
            fmt.Println("fizz")
        } else if i % 5 == 0 {
            fmt.Println("buzz")
        }
    }
}
