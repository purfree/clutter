package doubly

type List struct {
	head *Node
	tail *Node
	size int
}

func New() *List {
	l := &List{
		head: newNode(nil),
		tail: newNode(nil),
		size: 0,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (p *List) Length() int {
	return p.size
}

// 添加数据到链尾
func (p *List) Append(data interface{}) *Node {
	n := newNode(data)
	// 链尾节点
	t := p.tail.prev

	t.next = n
	n.prev = t

	n.next = p.tail
	p.tail.prev = n
	p.size++
	return n
}

// 添加数据到链头
func (p *List) InsertFront(data interface{}) *Node {
	n := newNode(data)
	h := p.head.next

	n.next = h
	h.prev = n

	n.prev = p.head
	p.head.next = n
	p.size++
	return n
}

func (p *List) Remove(node *Node) bool {
	if p.size == 0 {
		return false
	}

	if node == nil {
		return false
	}

	c := p.head
	for c.Next() != nil {
		if c.Next() == node {
			c.next = c.next.next
			c.next.prev = c
			p.size--
			return true
		}
		c = c.Next()
	}
	return false
}

func (p *List) Front() *Node {
	return p.head.next
}

func (p *List) Back() *Node {
	return p.tail.prev
}
