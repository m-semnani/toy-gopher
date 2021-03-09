package main

import "testing"

func TestMaxHeap_left(t *testing.T) {
	type fields struct {
		array []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "first",
			fields: struct{ array []int }{array: []int{
				1, 2, 3,
			}},
			args: struct{ i int }{i: 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MaxHeap{
				array: tt.fields.array,
			}
			if got := h.left(tt.args.i); got != tt.want {
				t.Errorf("left() = %v, want %v", got, tt.want)
			}
		})
	}
}
