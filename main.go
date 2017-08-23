package main

import (
	"fmt"
	"reflect"
)

// represents transaction with a hash of operations
type Node struct {
	parent *Node
	data   map[string]string
}

type Root struct {
	head *Node
}

// start of REPL: make first Node and set as head
func initialize() *Root {
	m := make(map[string]string)
	newNode := &Node{parent: nil, data: m}
	return &Root{head: newNode}
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
	// initialize program with a default Node and set to currentNode
	root := initialize()
	current := root.head
	fmt.Println(reflect.TypeOf(root))

	fmt.Println(len(current.data))
	current.Write("cat", "cat")
	fmt.Println(len(current.data))
	fmt.Println(current.Read("cat"))
}
