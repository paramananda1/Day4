package main

import (
	"fmt"
)

type item struct {
	value interface{}
	next  *item
}

type Stack struct {
	top  *item
	size int
}


// Stack method

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) Push(value interface{}) {
	fmt.Println("Push to stack ",value)
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *Stack) Pop() (value interface{}) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}

	return nil
}

// Queue Methods

type node struct {
	value interface{}
	next  *node
	prev  *node
}

type Queue struct {
	item  *node
	size int
}

func (Q *Queue) Len() int {
	return Q.size
}

func (Q *Queue) PushBack(value interface{}) {
	//fmt.Println("Push to Rear of Queue ",value)

	if Q.size == 0 {

		Q.item = &node{
			value: value,
			next: nil,
			prev: nil,
		}
		//fmt.Println("Push to Rear of Queue : ",value)
		//fmt.Println("Push to Rear of Queue : ",&Q.item,Q.item)
	}else{
		//fmt.Println("Push to Rear of Queue : ",&Q.item,Q.item)
		newnode := &node{
			value: value,
			next: nil,
			prev: nil,
		}
		tempnode := Q.item
		for tempnode.prev != nil {
			tempnode = tempnode.prev
		}
		newnode.next = tempnode
		tempnode.prev = newnode
		//Q.item=newnode
		//fmt.Println("Push to Rear of Queue : ",newnode,&newnode)
		//fmt.Println("Push to Rear of Queue : ",value)
	}
	fmt.Println("Push to Rear of Queue : ",value)
	Q.size++
}

func (Q *Queue) PopFront() (value interface{}) {
	//fmt.Println("Pop from front of Queue ",&Q.item,Q.item)

	if Q.Len()  > 0 {
		value = Q.item.value
		if Q.item.prev != nil{
			Q.item.prev.next = nil
			//Q.item.prev.next.prev = nil
			Q.item = Q.item.prev
		}
		Q.size--
		return
	}

	return nil
}

func StackMain() {


	stack := new(Stack)
	// Push different data type to the stack
	stack.Push(1)
	stack.Push("Welcome")
	stack.Push(4.0)
	stack.Push(3.4)

	// Pop until stack is empty
	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}


	Q := new(Queue)
	// Push different data type to the stack
	Q.PushBack(1)
	Q.PushBack("Welcome")
	Q.PushBack(4.0)
	Q.PushBack(3.4)

	// Pop until stack is empty
	for Q.Len() > 0 {
		fmt.Println("Pop from front of Queue : ",Q.PopFront())
	}
}
