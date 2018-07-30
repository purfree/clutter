package singly

import (
	"github.com/pkg/errors"
)

type Node struct {
	data interface{}
	next *Node
}

func (p *Node) Next() *Node {
	return p.next
}

func (p *Node) Data() interface{} {
	return p.data
}

type List struct {
	head *Node
	size int
}

func New() *List {
	return &List{}
}

func (p *List) Length() int {
	return p.size
}

// 添加数据到链尾
func (p *List) Append(data interface{}) *Node {
	n := &Node{}
	n.data = data
	if p.size == 0 {
		p.head = n
	} else {
		tem := p.head
		for tem != nil {
			if tem.next == nil {
				tem.next = n
				break
			}
			tem = tem.next
		}
	}
	p.size++
	return n
}

// 添加数据到指定位置
// 如果index大于等于链长度，添加到链尾
// 如果index小于0，添加到链头
func (p *List) Insert(data interface{}, index int) *Node {
	if index >= p.size {
		return p.Append(data)
	}
	if index < 0 {
		index = 0
	}
	n := &Node{}
	n.data = data
	if index == 0 {
		n.next = p.head
		p.head = n
	} else {
		tem := p.head
		for tem != nil {
			index--
			if index == 0 {
				n.next = tem.next
				tem.next = n
				break
			}
			tem = tem.next
		}
	}
	p.size++
	return n
}

func (p *List) Delete(node *Node) error {
	if p.size == 0 {
		return nil
	}
	if node == nil {
		return nil
	}
	if p.head == node {
		p.head = p.head.next
		p.size--
		return nil

	}
	tem := p.head
	for tem != nil {
		if tem.next == node {
			tem.next = tem.next.next
			p.size--
			return nil
		}
		tem = tem.next
	}
	return errors.Errorf("not found")
}

func (p *List) Remove(index int) error {
	if p.size == 0 {
		return nil
	}
	if index < 0 || index >= p.size {
		return errors.Errorf("out of range")
	}
	if index == 0 {
		p.head = p.head.next
		p.size--
		return nil
	}
	tem := p.head
	for tem != nil {
		index--
		if index == 0 {
			//if tem.next == nil {
			//	return errors.Errorf("out of range")
			//}
			tem.next = tem.next.next
			p.size--
			return nil
		} else if index < 0 {
			// 不可抵达
			panic("not arrived")
			//return errors.Errorf("not found")
		}
		tem = tem.next
	}
	return errors.Errorf("out of range")
}

func (p *List) Reverse() {
	if p.size <= 1 {
		return
	}
	if p.size == 2 {
		tem := p.head.next
		p.head.next = nil
		tem.next = p.head
		p.head = tem
		return
	}
	cur := p.head
	next := p.head.next
	cur.next = nil
	for next != nil {
		tem := next.next
		next.next = cur
		cur = next
		next = tem
	}
	p.head = cur
}
