package main

import "fmt"

func main() {
    go fizzBuzz()
}

func fizzBuzz() {
    for i := 0; i < 100 ; i++ {
        if i % (2 * 5) == 0 {
            fmt.Println(fmt.Sprint(i) + "FizzBuzz")
        } else if i % 5 == 0 {
            fmt.Println("Fizz")
        } else if i % 2 == 0 {
            fmt.Println("Buzz")
        }
    }
}
