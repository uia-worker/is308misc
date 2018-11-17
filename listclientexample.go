package main

import (
	"container/list"
	"fmt"
)

// List  an interface what should be implemented by classes
// Remember, - no explicit relation to between interface and class ...
type List interface {
	New() List
}

// ListClientExample encapsulates a list
type ListClientExample struct {
	list *list.List
}

// ListClientExample constructor
// "Pointer receiver" of type ListClientExample
func (lce *ListClientExample) ListClientExample() {
	lce.list = list.New()
	fmt.Printf("Type of lce.list is %T\n", lce.list)
	fmt.Printf("Value of lce.list is %v\n", lce.list)
}

func (lce *ListClientExample) getList() *list.List {
	return lce.list
}

func main() {

	lce := &ListClientExample{}
	lce.ListClientExample()
	l2 := lce.getList()
	fmt.Printf("Type of l2 is %T\n", l2)
	fmt.Printf("Value of l2 is %v\n", l2)
	l2.PushBack(4)
	fmt.Println(l2.Front().Value)

	// Create a new list and put some numbers in it.
	l := list.New()
	fmt.Printf("Type of l is %T\n", l)
	fmt.Printf("Value of l is %v\n", l)
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
