package main

import "fmt"

func main() {
    fizzBuzz()
}

func fizzBuzz() {
    for i := 0; i < 100 ; i++ {
        if i % (3 * 5) == 0 {
            fmt.Println(fmt.Sprint(i) + ": FizzBuzz")
        } else if i % 5 == 0 {
            fmt.Println(fmt.Sprint(i) + ": Fizz")
        } else if i % 3 == 0 {
            fmt.Println(fmt.Sprint(i) + ": Buzz")
        }
    }
}
