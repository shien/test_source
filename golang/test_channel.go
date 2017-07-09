package main

import "fmt"
import "strconv"
import "time"

func main() {

    c := make(chan int)
    go fizzBuzz(2, c)
    defer close(c)

    var i = <- c

    time.Sleep(5 * time.Millisecond)
    fmt.Println(strconv.Itoa(i))
}

func fizzBuzz(num int, s chan<- int) {

    for i := 0; i < 100 ; i++ {
        if i % (3 * 5) == 0 {
            fmt.Println(strconv.Itoa(i) + ": FizzBuzz -> " + strconv.Itoa(num))
        } else if i % 5 == 0 {
            fmt.Println(strconv.Itoa(i) + ": Fizz -> " + strconv.Itoa(num))
        } else if i % 3 == 0 {
            fmt.Println(strconv.Itoa(i) + ": Buzz -> " + strconv.Itoa(num))
        }
    }
    s <- num

}
