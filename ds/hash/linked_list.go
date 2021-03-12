package main

import (
	"fmt"
	"strconv"
)

type LinkedList struct {
	head *Node
	size int
}

type Node struct {
	data int
	next *Node
}

func (l *LinkedList) add(data int) {
	temp := l.head
	for {
		if temp.next == nil {
			tail := &Node{
				data: data,
				next: nil,
			}
			temp.next = tail
			l.size++
			return
		}
		temp = temp.next
	}
}

func (l *LinkedList) delByData(data int) bool {
	temp := l.head
	for temp.next != nil && temp.next.data != data {
		temp = temp.next
	}
	if temp.next == nil {
		fmt.Printf("%v not found to delete\n", data)
		return false
	} else {
		temp.next = temp.next.next
		l.size--
		fmt.Printf("%v deleted successfully!\n", data)
		return true
	}
}

func (l *LinkedList) insert(data int, place int) {
	if place < 0 || place > l.size {
		println("Index out of bound.")
		return
	}

	i := 0
	temp := l.head
	for {
		if i == place {
			node := &Node{data, temp.next}
			temp.next = node
			l.size++
			return
		}
		temp = temp.next
		i++
	}

}

func (l *LinkedList) traverse() string {
	temp := l.head
	result := "{ " + strconv.Itoa(temp.data)
	for temp.next != nil {
		temp = temp.next
		result = result + ", " + strconv.Itoa(temp.data)
	}
	result = result + " }"
	return result
}

func main() {
	ll := &LinkedList{&Node{20, nil}, 0}
	for i := 0; i < 20; i++ {
		ll.add(i)
	}
	println(ll.size)
	println(ll.traverse())

	ll.delByData(10)
	println(ll.traverse())

	ll.delByData(19)
	println(ll.traverse())

	ll.delByData(23)
	println(ll.traverse())
}
