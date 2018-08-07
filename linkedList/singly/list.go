package singly

type List struct {
	head *Node
	size int
}

func New() *List {
	return &List{
		head: newNode(nil),
		size: 0,
	}
}

func (p *List) Length() int {
	return p.size
}

// 添加数据到链尾
func (p *List) Append(data interface{}) *Node {
	n := newNode(data)
	c := p.head
	for c.next != nil {
		c = c.next
	}
	c.next = n
	p.size++
	return n
}

// 添加数据到指定位置
func (p *List) Insert(data interface{}, index int) *Node {
	if index < 0 || index > p.size-1 {
		return nil
	}
	n := newNode(data)
	c := p.head
	for c.next != nil {
		if index == 0 {
			break
		}
		index--
		c = c.next
	}
	n.next = c.next
	c.next = n
	p.size++
	return n
}

// 删除节点
func (p *List) Remove(node *Node) bool {
	if p.size == 0 {
		return false
	}
	if node == nil {
		return false
	}
	c := p.head
	for c.next != nil {
		if c.next == node {
			c.next = c.next.next
			p.size--
			return true
		}
		c = c.next
	}
	return false
}

// 删除指定位置的节点
func (p *List) RemoveAt(index int) bool {
	if index < 0 || index > p.size-1 {
		return false
	}
	c := p.head
	for c.next != nil {
		if index == 0 {
			break
		}
		index--
		c = c.next
	}
	c.next = c.next.next
	p.size--
	return true
}

// 获取指定位置的节点
func (p *List) Get(index int) *Node {
	if index < 0 || index > p.size-1 {
		return nil
	}
	c := p.head
	for c.next != nil {
		if index == 0 {
			break
		}
		index--
		c = c.next
	}
	return c.next
}

// 反转
func (p *List) Reverse() {
	if p.size <= 1 {
		return
	}
	pre := p.head.next
	cur := pre.next
	pre.next = nil

	// 迭代
	for cur != nil {
		next := cur.next
		cur.next = pre

		pre = cur
		cur = next
	}
	p.head.next = pre

	// 递归
	//p.head.next = nil
	//p.recursion(pre, cur)
}

func (p *List) recursion(pre, cur *Node) {
	if cur.next != nil {
		p.recursion(cur, cur.next)
	}
	cur.next = pre
	pre.next = nil
	if p.head.next == nil {
		p.head.next = cur
	}
}
