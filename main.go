package main

import (
	"fmt"
)

// represents transaction with a hash of operations
type Node struct {
	parent *Node
	data   map[string]string
}

func NewNode(parent *Node) *Node {
	m := make(map[string]string)
	return &Node{data: m}
}

// store k,v to current node's data
func (n *Node) Write(k, v string) string {
	d := n.data
	d[k] = v
	return ""
}

func (n *Node) Read(k string) string {
	d := n.data
	return d[k]
}

func main() {
	// initialize with new node as head
	head := NewNode(nil)

	fmt.Println(len(head.data))
	head.Write("cat", "cat")
	fmt.Println(len(head.data))
	fmt.Println(head.Read("cat"))
}
