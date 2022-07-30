package main

import "fmt"
import "structures/list"

func main() {
	head := list.New()
	fmt.Print("empty list:  ")
	head.Display() // empty list

	fmt.Print("insert at end: ")
	head.InsertAtEnd(5)
	head.InsertAtEnd(12)
	head.InsertAtEnd(25)
	head.Display()

	fmt.Print("insert at start: ")
	head.InsertAtStart(8)
	head.InsertAtStart(13)
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
