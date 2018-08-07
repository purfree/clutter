package doubly

import (
	"fmt"
	"testing"
)

func print(l *List) {
	if l.size == 0 {
		return
	}
	fmt.Printf("size: %v \n", l.Length())
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Printf("%v ", n.Data())
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

func TestList_List(t *testing.T) {
	//l := New()
	//l.Append(0)
	//l.Append(1)
	//fmt.Println(l)
	//l2 := glist.New()

	//l.PushBack(1)
	//l.PushBack(2)
	//for n := l.Front(); n != nil; n = n.Next() {
	//	fmt.Println(n.Value)
	//}
}

func TestList_Append(t *testing.T) {
	l := append()
	print(l)

	fmt.Println(l.Front().Data())
	fmt.Println(l.Back().Data())
}

func TestList_InsertFront(t *testing.T) {
	l := append()
	print(l)

	l.InsertFront(10)
	l.InsertFront(11)
	l.InsertFront(12)

	print(l)
}

func TestList_Remove(t *testing.T) {
	l := append()

	n := l.Append(10)
	print(l)
	l.Remove(n)
	print(l)

	n = l.InsertFront(10)
	print(l)
	l.Remove(n)
	print(l)

}
