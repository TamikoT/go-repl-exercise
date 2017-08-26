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

func Start(head *Transaction) *Transaction {
	// parent is current
	n := NewTransaction(head)
	// copy map data
	for k, v := range head.data {
		n.data[k] = v
	}
	return n
}

func isHead(t *Transaction) bool {
	if t.parent == nil {
		return true
	} else {
		return false
	}
}

func Abort(head *Transaction) *Transaction {
	head = head.parent
	// destroy current node?
	return head
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

func main() {
	// initialize with no parent
	head := NewTransaction(nil)
	fmt.Println(head) // to see if it's working...

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		fmt.Print("> ")
		input := strings.Split(scanner.Text(), " ")
		if input[0] == "quit" {
			break
		} else {
			switch strings.ToUpper(input[0]) {
			case "START":
				head = Start(head)
				fmt.Println(head) //to see if it's working...
			case "WRITE":
				head.Write(input[1], input[2])
			case "READ":
				if val, err := head.Read(input[1]); err == nil {
					fmt.Println(val)
				} else {
					fmt.Println(err)
				}
			case "ABORT":
				if isHead(head) == true {
					fmt.Println("ERROR: ABORT called with no active transaction.")
				} else {
					head = Abort(head)
					fmt.Println(head) // to see if it's working...
				}
			case "COMMIT":
				// TODO: change committed to collapst the last transactions into the last
			default:
				fmt.Println("ERROR: Unknown command: " + input[0])
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
