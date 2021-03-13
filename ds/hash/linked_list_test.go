package main

import "testing"

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
			"remove zero size ll",
			initLL([]int{}),
			args{index: 0},
			initLL([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initial.delByIndex(tt.args.index)

			if tt.want.size != tt.initial.size {
				t.Errorf("wanted size=%v but actual size was %v=", tt.want.size, tt.initial.size)
			}

			i1 := tt.want.head
			i2 := tt.initial.head
			for i1 != nil {
				if i1.data != i2.data {
					t.Errorf("want=%v but actual was %v=", i1.data, i2.data)
				}
			}

		})
	}
}
