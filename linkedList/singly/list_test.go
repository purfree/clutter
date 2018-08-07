package singly

import (
	"fmt"
	"testing"
)

func print(l *List) {
	if l.size == 0 {
		return
	}
	fmt.Printf("size: %v \n", l.Length())
	tem := l.head
	for tem != nil {
		fmt.Printf("%v ", tem.Data())
		tem = tem.next
	}
	fmt.Println()
}

func append() *List {
	l := New()
	for i := 0; i < 10; i++ {
		l.Append(i)
	}
	return l
}

func insert() *List {
	l := New()
	for i := 0; i < 10; i++ {
		l.Insert(i, 5)
	}
	return l
}

func TestList_Append(t *testing.T) {
	l := append()
	print(l)
}

func TestList_Insert(t *testing.T) {
	l := insert()
	print(l)

	l.Insert(22, 22)
	print(l)
}

func TestList_Delete(t *testing.T) {
	l := append()

	a := l.Append(22)
	print(l)
	l.Delete(a)
	print(l)

	b := l.Insert(33, 5)
	print(l)
	l.Delete(b)
	print(l)

	c := l.Insert(33, 0)
	print(l)
	l.Delete(c)
	print(l)
}

func TestList_Remove(t *testing.T) {
	l := append()

	print(l)
	l.Remove(0)
	print(l)

	print(l)
	l.Remove(8)
	print(l)

	print(l)
	l.Remove(33)
	print(l)

	print(l)
	l.Remove(3)
	print(l)
}

func TestList_Reverse(t *testing.T) {
	l := append()
	print(l)
	l.Reverse()
	print(l)
}

func TestList_Get(t *testing.T) {
	l := append()
	print(l)
	n := l.Get(0)
	fmt.Println(n.Data())
	n = l.Get(9)
	fmt.Println(n.Data())
}
