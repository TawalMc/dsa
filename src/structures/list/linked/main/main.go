package main

import (
	"fmt"
	"structures/list/linked"
)

func main() {
	head := linked.New()
	fmt.Print("empty list:  ")
	head.Display() // empty list

	fmt.Print("insert at end: ")
	head.InsertAtEnd(15)
	head.InsertAtEnd(23)
	head.InsertAtEnd(7)
	head.Display()

	fmt.Print("insert at start: ")
	head.InsertAtStart(88)
	head.InsertAtStart(163)
	head.Display()

	fmt.Print("insert at start: ")
	head.InsertAt(69, 0)
	head.Display()

	fmt.Print("delete data: ")
	head.DeleteAt(4)
	head.DeleteAt(1)
	head.Display()

	fmt.Print("reversed list: ")
	head.ReversedList()
	head.Display()
}
