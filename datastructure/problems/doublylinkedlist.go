package main


type Node struct{
	data int
	next *Node
	prev *Node
}

type LinkedList struct{
	head *Node
	tail *Node
}

func (l *LinkedList) addfront(data int)  {
	newNode := &Node{
		data: data,
		next: l.head,
		prev: nil,
	}

	if l.head == nil{
		l.head = newNode
		l.tail = newNode
		return
	}

	l.head.prev = newNode
	l.head = newNode
}

func (l *LinkedList) addback(data int)  {
	newNode := &Node{
		data :data,
		prev: l.tail,
		next: nil,
	
	}
	if l.tail == nil{
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	l.tail = newNode

}

func (l *LinkedList) display()  {
	
}

func main()  {
	list := LinkedList{}
	list.addback(50)
}