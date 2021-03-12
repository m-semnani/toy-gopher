package main

import (
	"fmt"
	"math/rand"
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

func main() {
	ll := &LinkedList{ &Node{12, nil}, 0}
	for i := 0; i < 20; i++ {
		ll.add(rand.Intn(20))
	}
	println(ll.size)
	fmt.Println(ll)
}
