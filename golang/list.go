package main

import (
	"fmt"
	"strconv"
)

type List struct {
	Data int
	Next *List
}

func (list *List) append(Data int) {
	var next *List = list
	for next.Next != nil {
		next = next.Next
	}
	next.Next = new(List)
	next.Next.Data = Data
}

func (list *List) show() {
	var next *List = list
	for next != nil {
		fmt.Println(next.Data)
		next = next.Next
	}
}

func (list *List) count(Data int) int {
	var next *List = list
	c := 0
	for next != nil {
		if next.Data == Data {
			c = c + 1
		}
		next = next.Next
	}
	return c
}

func main() {
	l := new(List)
	l.Data = 1
	l.Next = nil

	fmt.Println("append")
	l.append(2)
	l.append(3)
	l.append(3)
	l.append(4)
	l.append(5)
	l.append(2)
	l.append(7)
	l.append(8)
	l.append(3)
	fmt.Println("show")
	l.show()
	fmt.Println("count 3 is " + strconv.Itoa(l.count(3)))
}
