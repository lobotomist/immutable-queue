package queue

import "errors"

// ErrEmptyQueue error
var ErrEmptyQueue = errors.New("queue is an empty, enqueue or dequeue first")

type (
	// ImmutableQueue is an struct
	ImmutableQueue struct {
		back *ImmutableStack

		// we can cache the expanded linked list directly into an array and shift the pointer 'size'
		front []interface{}
		size  int
	}

	// ImmutableStack required for implementing queue
	// I choose linked-list structure for it
	ImmutableStack struct {
		v    interface{}
		next *ImmutableStack
	}
)

// NewImmutableQueue constructor
func NewImmutableQueue() *ImmutableQueue {
	return &ImmutableQueue{
		back: new(ImmutableStack),
	}
}

// EnQueue add new element to the queue
func (queue ImmutableQueue) EnQueue(v interface{}) *ImmutableQueue {
	return &ImmutableQueue{
		back:  queue.back.Push(v),
		front: queue.front,
		size:  queue.size,
	}
}

// DeQueue remove the element at the beginning of the immutable queue, and returns the new queue
func (queue *ImmutableQueue) DeQueue() *ImmutableQueue {
	if queue.size == 0 {
		var front []interface{}
		var back = queue.back

		for !back.IsEmpty() {
			front = append(front, back.Head())
			back = back.Pop()
		}

		return &ImmutableQueue{
			front: front,
			back:  back,
			size:  len(front) - 1,
		}
	}

	return &ImmutableQueue{
		front: queue.front,
		back:  queue.back,
		size:  queue.size - 1,
	}
}

// Head returns a first element or nil if queue is an empty
func (queue *ImmutableQueue) Head() (interface{}, error) {
	if queue.size == 0 {
		return nil, ErrEmptyQueue
	}
	return queue.front[queue.size-1], nil
}

// IsEmpty returns true an empty queue
func (queue *ImmutableQueue) IsEmpty() bool {
	return queue.size == 0
}

// Push new element to the stack
func (stack *ImmutableStack) Push(v interface{}) *ImmutableStack {
	return &ImmutableStack{
		v:    v,
		next: stack,
	}
}

// Pop returns next stack object
func (stack *ImmutableStack) Pop() *ImmutableStack {
	return stack.next
}

// Head returns the top element
func (stack *ImmutableStack) Head() interface{} {
	return stack.v
}

// IsEmpty retruns true if stack is an empty
func (stack *ImmutableStack) IsEmpty() bool {
	return stack == nil
}

// Reverse returns a copy of reversed stack
func (stack *ImmutableStack) Reverse() *ImmutableStack {
	copy := &ImmutableStack{}

	for !stack.IsEmpty() {
		copy = copy.Push(stack.v)
		stack = stack.Pop()
	}

	return copy
}
