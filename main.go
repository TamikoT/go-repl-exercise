package main

import (
	"fmt"
)

type Transaction struct {
	parent *Transaction
	data   map[string]string
}

func NewTransaction(p *Transaction) *Transaction {
	m := make(map[string]string)
	return &Transaction{parent: p, data: m}
}

func (n *Transaction) Write(k, v string) {
	n.data[k] = v
}

func (t *Transaction) Read(k string) string {
	if v, ok := t.data[k]; ok {
		return v
	}
	return "Key not found: " + k
}

func Start(head *Transaction) *Transaction {
	// parent is current
	n := NewTransaction(head)
	// copy map data
	for k, v := range head.data {
		n.data[k] = v
	}
	return n
}

func main() {
	// initialize with new no parent
	head := NewTransaction(nil)

	fmt.Println(head)
	fmt.Println(len(head.data))
	head.Write("cat", "cat")
	fmt.Println(len(head.data))
	fmt.Println(head.Read("cat"))
	fmt.Println(head.Read("bat"))
	head = Start(head)
	fmt.Println(head)
	fmt.Println(len(head.data))
}
