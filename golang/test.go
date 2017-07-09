package main

import "fmt"
import "strconv"
import "time"

func main() {
	go fizzBuzz(1)
	go fizzBuzz(2)
	go fizzBuzz(3)
	time.Sleep(5 * time.Millisecond)
}

func fizzBuzz(num int) {
	for i := 0; i < 100; i++ {
		if i%(3*5) == 0 {
			fmt.Println(strconv.Itoa(i) + ": FizzBuzz -> " + strconv.Itoa(num))
		} else if i%5 == 0 {
			fmt.Println(strconv.Itoa(i) + ": Fizz -> " + strconv.Itoa(num))
		} else if i%3 == 0 {
			fmt.Println(strconv.Itoa(i) + ": Buzz -> " + strconv.Itoa(num))
		}
	}
}
