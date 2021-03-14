package main

import (
	"testing"
)

func initLL(d []int) *LinkedList {
	ll := &LinkedList{}
	for _, data := range d {
		ll.add(data)
	}
	return ll
}

func TestLinkedList_delByIndex(t *testing.T) {
	type fields struct {
		nodes []Node
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		initial *LinkedList
		args   args
		want *LinkedList
	}{
		{
			"remove from zero size ll",
			initLL([]int{}),
			args{index: 0},
			initLL([]int{}),
		},
		{
			"remove a tail",
			initLL([]int{1, 2, 3, 4, 5}),
			args{index: 5},
			initLL([]int{1, 2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initial.delByIndex(tt.args.index)

			if tt.want.size != tt.initial.size {
				t.Errorf("wanted size=%v but actual size was %v=", tt.want.size, tt.initial.size)
			}

			if tt.want.size != tt.initial.size {
				t.Errorf("wanted size=%v but actual was %v=", tt.want.size, tt.initial.size)
			}

			tempInit := tt.initial.head
			tempWant := tt.initial.head
			for tempInit != nil && tempWant !=nil  {
				if tempInit.data != tempWant.data {
					t.Errorf("Wanted data=%v but actual was %v=", tempWant.data, tempInit.data)
				}
				tempInit = tempInit.next
				tempWant = tempWant.next
			}
		})
	}
}
