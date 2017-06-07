package main

import "fmt"
import "strconv"

func main() {
    fizzBuzz()
}

func fizzBuzz() {
    for i := 0; i < 100 ; i++ {
        if i % (3 * 5) == 0 {
            fmt.Println(strconv.Itoa(i) + ": FizzBuzz")
        } else if i % 5 == 0 {
            fmt.Println(strconv.Itoa(i) + ": Fizz")
        } else if i % 3 == 0 {
            fmt.Println(strconv.Itoa(i) + ": Buzz")
        }
    }
}
