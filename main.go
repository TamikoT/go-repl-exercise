package main

import (
	"bufio"
	"fmt"
	"os"
)

// Data structure = linked list as a stack [data+head]->[2]->[1]->[0]

// acts as head of list & contains DB
type HeadDB struct {
	// access to stack of operations
	current *Node
	// most recent ops for same key
	data map[string]string
}

// represents transaction with a hash of operations
type Node struct {
	parent *Node
	ops    map[string]struct
}

type Read struct {
}

type Write struct {
	value string
}

type Delete struct {
}

// might need an interface for operations (?)

func (h *HeadDB) WriteTo(k,v string) *HeadDB {
	// reference new Write Struct
	writeOp := Write{value: v}
	// access current ops list
	list = &h.current.ops
	*list[k] = writeOp
}

func (h *HeadDB) Start() *HeadDB {
	newOps := make(map[string]Op)
	// newNode where parent is current
	newNode := &Node{parent: h.current ops: newOps }
	// Then reset Head's current to new node
	h.current = newNode
	return h
}

func (h *HeadDB) Abort() *HeadDB {
	// TODO: move the head pointer and delete current Node
	return h
}

func (h *HeadDB) Commit() *HeadDB {
	// TODO: write current commands to datastore
	return h
}


func main() {
	// initialize program with a default Node instance and set to currentNode
	// firstNode := Node{parent: nil}
	// master := HeadDB{}

	// loop for until input = 'quit'
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Start of program.")
	input, _ := reader.ReadString('\n')
	// WIP: switch on string
	fmt.Println(input)
	switch input {
	case "start":
		fmt.Println("matched start")
	case "write":
		fmt.Println("matched write")
	}
}

