package main

import "fmt"

type Tree struct {
	Left *Tree
	Right *Tree
	val string
}

func Preorder(head *Tree){
	if head == nil {
		return
	}
	fmt.Printf(head.val)
	Preorder(head.Left)
	Preorder(head.Right)
}

func Postorder(head *Tree){
	if head == nil {
		return
	}
	Postorder(head.Left)
	Postorder(head.Right)
	fmt.Printf(head.val)
}

func main() {
	head := &Tree{val:"+"}
	a := &Tree{val: "a"}
	minus := &Tree{val: "-"}
	b := &Tree{val: "b"}
	c := &Tree{val: "c"}

	head.Left = a
	head.Right = minus
	minus.Left = b
	minus.Right = c

	fmt.Println("Preorder: ")
	Preorder(head)
	fmt.Println()
	fmt.Println("Postorder: ")
	Postorder(head)
	fmt.Println()
}
