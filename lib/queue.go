package lib

//Node defines ll node
type Node struct {
	Value interface{}
	Next  *Node
}

//Queue defines ll queue
type Queue struct {
	Name string
	Tail *Node
	Head *Node
}

//CreateQueue instantiates a new queue
func CreateQueue(name string) *Queue {
	return &Queue{Name: name}
}

//Enqueue adds a new member to the queue
func (queue *Queue) Enqueue(value string) {
	node := &Node{Value: value, Next: nil}
	if queue.Head == nil {
		queue.Head = node
		return
	}
	queue.Tail = node
}

//Dequeue removes a member from the queue
func (queue *Queue) Dequeue(value string) *Node {
	if queue.Head == nil {
		return nil
	}
	node := queue.Head
	newHead := queue.Head.Next
	queue.Head = newHead
	return node
}
