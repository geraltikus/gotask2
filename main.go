package main

import "fmt"

type Node struct {
	ID        string
	Relations []Node
}

type AnotherNode struct {
	ID        string
	Relations []AnotherNode
}

func main() {
	a := Node{
		ID:        "a",
		Relations: nil,
	}

	b := Node{
		ID:        "b",
		Relations: nil,
	}

	c := Node{
		ID:        "c",
		Relations: []Node{a, b},
	}

	a.Relations = []Node{b, c}
	b.Relations = []Node{a, c}

	//	A --- B
	//	 \   /
	//	   C
	fmt.Printf("a: %+v", a)
}
