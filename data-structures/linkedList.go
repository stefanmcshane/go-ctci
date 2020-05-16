package ctci

type LinkedListNode struct {
	Data interface{}
	Next *LinkedListNode
	Prev *LinkedListNode
}

func (l *LinkedListNode) Append(d interface{}) {
	newNode := LinkedListNode{
		Data: d,
	}
	for {
		if l.Next == nil {
			break
		}
		l = l.Next
	}

	if l.Next != nil {
		panic("Data unable to find end of linkedlist")
	}

	newNode.Prev = l
	l.Next = &newNode
}

func (l *LinkedListNode) Remove(d interface{}) {

	// If single linked list, this is required, otherwise can use l.Prev
	prev := l
	for l.Next != nil {
		if l.Data == d {
			break
		}
		prev = l
		l = l.Next
	}
	if prev == nil {
		panic("Data not found in linkedlist")
	}
	prev.Next = l.Next
	l.Next.Prev = prev
}

func (l *LinkedListNode) Find(d interface{}) int {
	count := 0
	found := false

	for {
		if l.Data == d {
			found = true
			break
		}
		if l.Next == nil {
			break
		}
		l = l.Next
		count++
	}

	if found {
		return count
	}
	return -1
}
