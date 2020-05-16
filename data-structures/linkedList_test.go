package ctci

import (
	"testing"
)

var ll *LinkedListNode

func linkedListInit() {
	ll1 := &LinkedListNode{
		Data: 1,
	}
	ll2 := &LinkedListNode{
		Data: 2,
		Prev: ll1,
	}
	ll3 := &LinkedListNode{
		Data: 3,
		Prev: ll2,
	}
	ll2.Next = ll3
	ll1.Next = ll2

	ll = ll1
}

func TestLinkedListNode_Append(t *testing.T) {
	linkedListInit()
	type fields struct {
		Data interface{}
		Next *LinkedListNode
		Prev *LinkedListNode
	}
	type args struct {
		d interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "LinkedList-Append",
			fields: fields{
				Data: "test",
			},
			args: args{
				d: "New data",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedListNode{
				Data: tt.fields.Data,
				Next: tt.fields.Next,
				Prev: tt.fields.Prev,
			}
			l.Append(tt.args.d)
			if l != l.Next.Prev {
				t.Errorf("Linked List Append error, %+v %+v", l, l.Next.Prev)
			}
		})
	}
}

func TestLinkedListNode_Remove(t *testing.T) {
	linkedListInit()
	ll3 := &LinkedListNode{
		Data: 3,
	}

	type args struct {
		d interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "LinkedList-Remove",
			args: args{
				2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll.Remove(tt.args.d)

			if ll.Next.Data != ll3.Data {
				t.Errorf("Linked List Remove error, %+v %+v", ll.Next, ll3)
			}
		})
	}
}

func TestLinkedListNode_Find(t *testing.T) {
	linkedListInit()
	type args struct {
		d interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "LinkedList - Find 1",
			args: args{
				d: 1,
			},
			want: 0,
		},
		{
			name: "LinkedList - Find 3",
			args: args{
				d: 3,
			},
			want: 2,
		},
		{
			name: "LinkedList - Find 3",
			args: args{
				d: 4,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ll.Find(tt.args.d); got != tt.want {
				t.Errorf("LinkedListNode.Find() = %v, want %v", got, tt.want)
			} else {
				t.Log(got, tt.want)
			}
		})
	}
}
