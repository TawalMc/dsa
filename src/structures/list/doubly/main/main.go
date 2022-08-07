package main

import (
	"fmt"
	"structures/list/doubly"
)

func main() {
	head := doubly.New()
	fmt.Print("empty list:  ")
	head.Display() // empty list

	fmt.Print("insert at start: ")
	head.InsertAtStart(27)
	head.InsertAtStart(45)
	head.InsertAtStart(12)
	head.InsertAtStart(148)
	head.Display()

	fmt.Print("insert at end: ")
	head.InsertAtEnd(13)
	head.InsertAtEnd(55)
	head.InsertAtEnd(0)
	head.Display()

	fmt.Print("reversed print: ")
	head.ReversePrint()

	fmt.Print("insert at index: ")
	head.InsertAt(222, 1)
	head.InsertAt(98, 5)
	head.Display()
}
