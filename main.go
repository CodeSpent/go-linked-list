package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) addNode(n *Node) {
	if l.head == nil {
		l.head = n
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}
}

func (l *LinkedList) getNode(d int) (*Node, error) {
	current := l.head

	if current == nil {
		return nil, errors.New("list is empty")
	}

	if current.data == d {
		return current, nil
	}

	for current.next != nil {
		if current.data == d {
			return current, nil
		}
		break
	}

	return nil, errors.New("node not found")
}

func (l *LinkedList) removeNode(n *Node) {
	current := l.head
	previous := l.head

	if current == l.head {
		l.head = current.next
	} else {
		previous.next = current.next
	}

	previous = current
	current = current.next
}

func (l *LinkedList) insertNode(data int, pos int) error {
	if pos < 0 {
		return errors.New("position cannot be negative")
	}

	node := &Node{data: data}

	if pos == 0 {
		node.next = l.head
		l.head = node
		return nil
	}

	current := l.head
	for i := 0; i < pos-1 && current != nil; i++ {
		current = current.next
	}

	if current == nil {
		return errors.New("position out of range")
	}

	node.next = current.next
	current.next = node

	return nil
}

func (l *LinkedList) length() int {
	count := 0
	current := l.head

	for current != nil {
		count++
		current = current.next
	}

	return count
}

func (l *LinkedList) clear() {
	l.head = nil
}

func (l *LinkedList) print() {
	var sb strings.Builder
	current := l.head
	for current != nil {
		sb.WriteString(fmt.Sprintf("%d", current.data))
		if current.next != nil {
			sb.WriteString(" -> ")
		}
		current = current.next
	}
	sb.WriteString(" -> nil")

	fmt.Println(sb.String())
}

func (l *LinkedList) reverse() {
	var prev *Node
	current := l.head
	var next *Node

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

func presentOptions() (int, error) {
	fmt.Println("")
	fmt.Println("You have the following options:")
	fmt.Println("1.) Add a node to the list.")
	fmt.Println("2.) Find a node in the list.")
	fmt.Println("3.) Remove a node from the list.")
	fmt.Println("4.) Insert a node at a specific index in the list.")
	fmt.Println("5.) Display current state of the list.")
	fmt.Println("6.) Quit.")
	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return strconv.Atoi(text)

}

func main() {

	list := LinkedList{}

	for {
		option, err := presentOptions()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch option {
		case 1:
			fmt.Println("What data should be inserted?")

			var data int
			_, err := fmt.Scanf("%d", &data)
			if err != nil {
				fmt.Println("Error:", err)
			}

			list.addNode(&Node{data: data})
			list.print()
			break
		case 2:
			fmt.Println("What data should be searched?")

			var data int
			_, err := fmt.Scanf("%d", &data)
			if err != nil {
				fmt.Println("Error:", err)
			}

			node, err := list.getNode(data)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			fmt.Println("Node found:", node.data)
		case 3:
			fmt.Println("What data should be removed?")

			var data int
			_, err := fmt.Scanf("%d", &data)
			if err != nil {
				fmt.Println("Error:", err)
			}

			node, err := list.getNode(data)
			if err != nil {
				fmt.Println("Error:", err)
			}

			list.removeNode(node)
			list.print()
			break
		case 4:
			fmt.Println("What data should be inserted?")

			var data int
			_, err := fmt.Scanf("%d", &data)
			if err != nil {
				fmt.Println("Error:", err)
			}

			fmt.Println("What position should this data be inserted?")

			var pos int
			_, err = fmt.Scanf("%d", &pos)
			if err != nil {
				fmt.Println("Error:", err)
			}

			err = list.insertNode(data, pos)
			if err != nil {
				fmt.Println("Error:", err)
			}

			list.print()
			break
		case 5:
			list.print()
			break
		case 6:
			os.Exit(0)
		}
	}
}
