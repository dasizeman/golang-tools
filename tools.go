package tools

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func init() {
	// Seed the random numbers from the clock
	rand.Seed(time.Now().UnixNano())
}

// IntMax returns the larger of the two provided integers
func IntMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

// IntMin returns the smaller of the two provided integers
func IntMin(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

// RandomInt returns a random integer in the range [min, max]
func RandomInt(min, max int) int {
	return rand.Intn(max+1-min) + min
}

// ReadFileToStrings reads the lines in a text file to
// a slice of strings.  It does not check if the input
// is actually a text file
func ReadFileToStrings(path string) ([]string, error) {
	var lines []string
	inFile, err := os.Open(path)
	if err != nil {
		return lines, err
	}

	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, err
}

// Simple "generic" stack
// -----------------------------------------

// Stack is a stack that can hold any value, but
// you have to type assert yourself
type Stack struct {
	data   []interface{}
	topIdx int
}

// Peek returns the item at the top of the stack without popping it
func (stack *Stack) Peek() interface{} {
	return stack.data[stack.topIdx-1]
}

// Pop pops the top item off the stack
func (stack *Stack) Pop() interface{} {
	value := stack.Peek()
	stack.topIdx--
	return value
}

// Push pushes a new item to the top of the stack
func (stack *Stack) Push(value interface{}) {
	stack.topIdx++
	if stack.topIdx > len(stack.data) {
		stack.data = append(stack.data, value)
	} else {
		stack.data[stack.topIdx-1] = value
	}
}

// IsEmpty returns true if the stack is empty
func (stack *Stack) IsEmpty() bool {
	return stack.topIdx <= 0
}

// Height returns the number of items on the stack
func (stack *Stack) Height() int {
	return stack.topIdx
}

// StackQueue is a queue made from two stacks
type StackQueue struct {
	enqueueStack, dequeueStack Stack
}

// Enqueue adds an item to the queue
func (queue *StackQueue) Enqueue(item interface{}) {
	queue.enqueueStack.Push(item)
}

// Dequeue removes an item from the queue
func (queue *StackQueue) Dequeue() interface{} {
	if queue.dequeueStack.IsEmpty() &&
		!queue.enqueueStack.IsEmpty() {
		queue.swapStacks()
	}

	if queue.IsEmpty() {
		return nil
	}

	return queue.dequeueStack.Pop()
}

func (queue *StackQueue) swapStacks() {
	for !queue.enqueueStack.IsEmpty() {
		queue.dequeueStack.Push(queue.enqueueStack.Pop())
	}
}

// IsEmpty returns true if the queue is empty
func (queue *StackQueue) IsEmpty() bool {
	return queue.dequeueStack.IsEmpty() &&
		queue.enqueueStack.IsEmpty()
}

// Length returns the number of items in the queue
func (queue *StackQueue) Length() int {
	return queue.enqueueStack.Height() +
		queue.dequeueStack.Height()
}
