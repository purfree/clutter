package circular

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
		if n.Data() == nil {
			break
		}
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

func TestList_Append(t *testing.T) {
	l := append()
	print(l)

	fmt.Println(l.Front().Data())
	fmt.Println(l.Back().Data())
}
