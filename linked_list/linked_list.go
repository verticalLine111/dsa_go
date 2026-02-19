package main

type Node struct {
	value int
	next *Node
}
type LinkedList struct {
	size int
	head *Node
}

func (l *LinkedList) pushFront(val int)  {
	node := &Node{
		value: val,
		next: l.head,
	}
	l.head = node
	l.size++
}

func (l *LinkedList)pushBack(val int)  {
	node := &Node{
		value: val,
		next: nil,
	}
	if l.head == nil {
		l.head = node
		l.size++
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	
	current.next = node
	l.size++
}
func (l *LinkedList)popFront() (int, error) {
	if l.head == nil { 
		panic("list is empty")
	}
	value := l.head.value	
	l.head = l.head.next
	l.size--
	return value, nil
}
func (l *LinkedList)popBack() (int, error) {
	if l.head == nil {
		panic("list is empty")
	}
	if l.head.next == nil {
		value := l.head.value
		l.head = nil
		l.size--
		return value , nil
	}
	oneBeforeLast := l.head
	for oneBeforeLast.next.next != nil {
		oneBeforeLast = oneBeforeLast.next
	}
	value := oneBeforeLast.next.value
	oneBeforeLast.next = nil
	l.size--
	return value, nil
}

func (l *LinkedList)insertAt(index int, val int) error {
	if index < 0 || l.size < index {
		panic("out of range")
	}
	if index == 0 {
		node := &Node{
			value: val,
			next: l.head,
		}
		l.head = node
		l.size++
		return nil
	}
	current := l.head
	for i := 0; i < index - 1; i++ {
		current = current.next
	}
	newNode := &Node{
		value: val,
		next: current.next,
	}
	current.next = newNode
	l.size++
	return nil
}
func (l *LinkedList)removeAt(index int) error {
	if index < 0 || l.size <= index {
		panic("out of range")
	}
	if index == 0 { 
		l.head = l.head.next
		l.size--
		return nil
	}
	current := l.head
	for i := 0; i < index - 1; i++ {
		current = current.next
	}
	current.next = current.next.next
	l.size--
	return nil
}

func (l *LinkedList)get(index int) (int, error) {
	if index < 0 || l.size <= index {
		panic("out of range")
	}
	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value, nil
}
func (l *LinkedList)reverse()  {
	var prev *Node= nil
	current := l.head
	for current != nil {
		next := current.next	
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}
func (l *LinkedList) Size() int {
	return l.size
}

func main()  {
	linkedlist := &LinkedList{
		head: &Node{
			value: 1,
			next: nil,
		},
		size: 1,
	}
	linkedlist.pushFront(5)


	current := linkedlist.head
	for current != nil { 
		println(current.value)
		current = current.next
	}
}