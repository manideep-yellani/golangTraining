package main

import "fmt"

type node struct {
	next *node
	val  int
}

func main() {
	//h := node{val: 1}
	//head := &h,
	var head *node

	head = add(head, 10)
	add(head, 9)
	add(head, 34)
	//add(head, 3)
	//add(head, 4)
	//add(head, 33)
	dis(head)
	head = del(head, 5)

	dis(head)

}

func add(head *node, va int) *node {
	t := node{}
	t.val = va
	if head == nil {
		head = &t

		return &t
	}
	cur := head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &t
	return head
}

func dis(head *node) {
	cur := head
	for cur != nil {
		fmt.Printf(" %v", cur.val)

		cur = cur.next
	}
	fmt.Println()
}

func del(head *node, p int) *node {
	cur := head
	i := 1
	if p == 1 {
		return head.next
	}
	for i < p-1 && cur.next != nil {
		cur = cur.next
		i++
	}
	if cur.next == nil {
		fmt.Println("give correct number")
		return head
	}
	if i != p-1 {
		fmt.Println("think and give the correct position")
		return head
	}

	cur.next = cur.next.next
	return head
}
