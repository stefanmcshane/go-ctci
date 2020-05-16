package ctci

import (
	"testing"
)

var s Stack

func stackInit() {
	d1 := &StackNode{
		Data: 1,
	}
	d2 := &StackNode{
		Data: 2,
	}
	d3 := &StackNode{
		Data: 3,
	}
	d1.Next = d2
	d2.Next = d3
	s = Stack{
		Top: d1,
	}
}
func TestStack_Push(t *testing.T) {
	stackInit()
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want Stack
	}{
		{
			name: "Stack push",
			args: args{i: 1},
			want: s,
		},
	}
	for _, tt := range tests {
		s.Top = s.Top.Next
		t.Run(tt.name, func(t *testing.T) {
			s.Push(tt.args.i)
			if s.Top.Data != tt.args.i {
				t.Errorf("TestStack_Push() Error => got %+v wanted %+v", s, tt.want)
			}
		})
	}
}
func TestStack_Peek(t *testing.T) {
	stackInit()
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Stack peek - 1",
			want: 1,
		},
		{
			name: "Stack peek - 2",
			want: 2,
		},
		{
			name: "Stack peek - 3",
			want: 3,
		},
		{
			name: "Stack peek - nil",
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sn := s.Peek()
			if sn != nil && tt.want >= 0 {
				if sn.Data != tt.want {
					t.Errorf("TestStack_Peek() Error => got %+v wanted %+v", sn.Data, tt.want)
				}
				s.Top = s.Top.Next
			} else {
				t.Log("TestStack_Peek() => stack is empty")
			}
		})
	}
}
func TestStack_Pop(t *testing.T) {
	stackInit()
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Stack peek - 1",
			want: 1,
		},
		{
			name: "Stack peek - 2",
			want: 2,
		},
		{
			name: "Stack peek - 3",
			want: 3,
		},
		{
			name: "Stack peek - nil",
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sn := s.Pop()
			if sn != nil && tt.want >= 0 {
				if sn.Data != tt.want {
					t.Errorf("TestStack_Peek() Error => got %+v wanted %+v", sn.Data, tt.want)
				}
			} else {
				t.Log("TestStack_Peek() => stack is empty as expected")
			}
		})
	}
}
