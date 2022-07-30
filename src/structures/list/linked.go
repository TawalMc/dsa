package list

import "fmt"

// Node is the type of each linked list child
type Node struct {
	data int
	next *Node
}

// New will create a head for a new linked list
func New() *Node {
	return &Node{0, nil}
}

// Traverse will go to the linked list's end
// "H" of receivers means "Head"
func (H *Node) Traverse() *Node {
	temp := H
	for temp.next != nil {
		temp = temp.next
	}

	return temp
}

// Display will output the linked list
func (H *Node) Display() {
	temp := H.next
	len := 0
	for temp != nil {
		len++
		fmt.Printf("%v ", temp.data)
		temp = temp.next
	}
	fmt.Printf("(len = %v)\n", len)
}

// InsertAtEnd will insert a new node at the end of the linked list
func (H *Node) InsertAtEnd(data int) {
	newNode := &Node{data, nil}
	tail := H.Traverse()

	tail.next = newNode
}

// InsertAtStart will insert a new node at the beginning of the linked list
func (H *Node) InsertAtStart(data int) {
	newNode := &Node{data, nil}

	newNode.next = H.next
	H.next = newNode
}

// InsertAt will insert data at any position
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
	for i := 0; temp.next != nil; i++ {
		if i == nth {
			newNode := &Node{data, nil}

			newNode.next = temp.next
			temp.next = newNode
			return
		}
		temp = temp.next
	}
}

func (H *Node) DeleteAt(nth int) {
	temp := H
	for i := 0; temp.next != nil; i++ {
		if i == nth {
			// temp1 := temp.next
			// temp.next = temp1.next
			// temp1 = nil
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
}

// TODO
func (H *Node) ReversedList() {
	temp := H
	for temp.next != nil {
		copyTemp := temp

		copyTemp.next = H.next
		H.next = copyTemp
		temp = temp.next
	}
}
