package circular

type List struct {
	root Node
	size int
}

func New() *List {
	l := new(List)
	l.root.next = &l.root
	l.root.prev = &l.root
	return l
}

func (p *List) Length() int {
	return p.size
}

func (p *List) Front() *Node {
	return p.root.next
}

func (p *List) Back() *Node {
	return p.root.prev
}

// 添加数据到链尾
func (p *List) Append(data interface{}) *Node {
	node := newNode(data)

	t := p.root.prev
	n := t.next

	t.next = node
	node.prev = t

	node.next = n
	n.prev = node
	p.size++
	return n
}
