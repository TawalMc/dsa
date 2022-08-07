package doubly

import "fmt"

type Head struct {
	Next *Node
}

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

func New() *Head {
	return &Head{nil}
}

func (H *Head) Display() {
	current := H.Next
	length := 0
	for current != nil {
		length++
		fmt.Printf("%v ", current.Data)
		current = current.Next
	}
	fmt.Printf("(length = %v)\n", length)
}

func (H *Head) InsertAtEnd(data int) {
	current := H.Next
	var prev *Node = current
	for current != nil {
		prev = current
		current = current.Next
	}
	if prev == nil {
		H.InsertAtStart(data)
	} else {
		newNode := &Node{data, nil, nil}
		prev.Next = newNode
		newNode.Prev = prev
	}
}

func (H *Head) InsertAtStart(data int) {
	newNode := &Node{data, nil, nil}
	current := H.Next

	if current == nil {
		H.Next = newNode
	} else {
		current.Prev = newNode
		newNode.Next = current
		H.Next = newNode
	}
}

func (H *Head) InsertAt(data, nth int) {
	current := H.Next

	if nth == 0 || current == nil {
		H.InsertAtStart(data)
		return
	}
	if nth == -1 {
		H.InsertAtEnd(data)
		return
	}

	for i := 0; current != nil; i++ {
		newNode := &Node{data, nil, nil}
		if i == nth {
			current.Prev.Next = newNode
			newNode.Next = current
			newNode.Prev = current.Prev
			current.Prev = newNode
		}
		current = current.Next
	}
}

func (H *Head) ReversePrint() {
	current := H.Next
	prev := current
	for current != nil {
		prev = current
		current = current.Next
	}
	length := 0
	for prev != nil {
		length++
		fmt.Printf("%v ", prev.Data)
		prev = prev.Prev
	}
	fmt.Printf("(length = %v)\n", length)
}
