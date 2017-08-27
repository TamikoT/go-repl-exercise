package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// transaction : represented as node of linked list
type transaction struct {
	parent *transaction
	data   map[string]string
}

// start : creates  new transaction
func start(parent *transaction) *transaction {
	newData := make(map[string]string)
	head := &transaction{parent: parent, data: newData}
	// if transaction is not the tail,copy data
	if !(parent == nil) {
		for k, v := range parent.data {
			head.data[k] = v
		}
	}
	return head
}

// abort : sets head to previous transaction
func abort(head *transaction) *transaction {
	return head.parent
}

// commit : deletes parent + points current node to parent's parent
func commit(head *transaction) *transaction {
	prevParent := head.parent.parent
	head.parent = prevParent
	return head
}

// isTail : checks if current node is tail
func isTail(current *transaction) bool {
	if current.parent == nil {
		return true
	}
	return false
}

func (t *transaction) write(k, v string) {
	t.data[k] = v
}

func (t *transaction) read(k string) (string, error) {
	if v, ok := t.data[k]; ok {
		return v, nil
	}
	return "", fmt.Errorf("Key not found: %s", k)
}

func (t *transaction) delete(k string) error {
	if _, ok := t.data[k]; ok {
		delete(t.data, k)
		fmt.Println("Key deleted: " + k)
		return nil
	}
	return fmt.Errorf("Key not found: %s", k)
}

func main() {
	head := start(nil)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		fmt.Print("> ")
		input := strings.Split(scanner.Text(), " ")

		if input[0] == "quit" {
			break
		}

		switch strings.ToUpper(input[0]) {
		case "START":
			head = start(head)
		case "WRITE":
			head.write(input[1], input[2])
		case "READ":
			if val, err := head.read(input[1]); err == nil {
				fmt.Println(val)
			} else {
				fmt.Println(err)
			}
		case "DELETE":
			head.delete(input[1])
		case "ABORT":
			if isTail(head) == true {
				fmt.Println("ERROR: ABORT called with no active transaction.")
			} else {
				head = abort(head)
			}
		case "COMMIT":
			head = commit(head)
			fmt.Println(head)
		default:
			fmt.Println("ERROR: Unknown command: " + input[0])
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
