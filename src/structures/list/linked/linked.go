package linked

import "fmt"

// Node is the type of each linked list child
type Node struct {
	Data int
	Next *Node
}

// New will create a head for a new linked list
func New() *Node {
	return &Node{0, nil}
}

// Traverse will go to the linked list's end
// "H" of receivers means "Head"
func (H *Node) Traverse() *Node {
	temp := H
	for temp.Next != nil {
		temp = temp.Next
	}

	return temp
}

// Display will output the linked list
func (H *Node) Display() {
	temp := H.Next
	length := 0
	for temp != nil {
		length++
		fmt.Printf("%v ", temp.Data)
		temp = temp.Next
	}
	fmt.Printf("(length = %v)\n", length)
}

// InsertAtEnd will insert a new node at the end of the linked list
func (H *Node) InsertAtEnd(data int) {
	newNode := &Node{data, nil}
	tail := H.Traverse()

	tail.Next = newNode
}

// InsertAtStart will insert a new node at the beginning of the linked list
func (H *Node) InsertAtStart(data int) {
	newNode := &Node{data, nil}

	newNode.Next = H.Next
	H.Next = newNode
}

// InsertAt will insert Data at any position
func (H *Node) InsertAt(data int, nth int) {

	if nth == 0 {
		H.InsertAtStart(data)
		return
	}
	if nth == -1 {
		H.InsertAtEnd(data)
		return
	}

	temp := H
	for i := 0; temp.Next != nil; i++ {
		if i == nth {
			newNode := &Node{data, nil}

			newNode.Next = temp.Next
			temp.Next = newNode
			return
		}
		temp = temp.Next
	}
}

func (H *Node) DeleteAt(nth int) {
	temp := H
	for i := 0; temp.Next != nil; i++ {
		if i == nth {
			// temp1 := temp.Next
			// temp.Next = temp1.Next
			// temp1 = nil
			temp.Next = temp.Next.Next
			return
		}
		temp = temp.Next
	}
}

func (H *Node) ReversedList() {
	var (
		current       = H.Next
		prev    *Node = nil
		next    *Node = nil
	)

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	H.Next = prev
}
