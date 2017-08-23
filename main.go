package main

import (
	"fmt"
)

type Transaction struct {
	parent *Transaction
	data   map[string]string
}

func NewTransaction(parent *Transaction) *Transaction {
	m := make(map[string]string)
	return &Transaction{data: m}
}

// store k,v to current data
func (n *Transaction) Write(k, v string) {
	d := n.data
	d[k] = v
}

func (n *Transaction) Read(k string) string {
	d := n.data
	return d[k]
}

func main() {
	// initialize with new no parent
	head := NewTransaction(nil)

	fmt.Println(len(head.data))
	head.Write("cat", "cat")
	fmt.Println(len(head.data))
	fmt.Println(head.Read("cat"))
}
