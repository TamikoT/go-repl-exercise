package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func (t *Transaction) Read(k string) (string, error) {
	if v, ok := t.data[k]; ok {
		return v, nil
	}
	return "", fmt.Errorf("Key not found: %s", k)
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

	head = Start(head)
	fmt.Println(head)
	fmt.Println(len(head.data))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		if input[0] == "quit" {
			break
		} else {
			switch input[0] {
			case "start":
				fmt.Println("Start called")
				Start(head)
			case "write":
				head.Write(input[1], input[2])
				fmt.Println(head.data) //to see if it's working...
			case "read":
				fmt.Println(head.Read(input[1]))
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
