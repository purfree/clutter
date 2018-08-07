package doubly

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

func newNode(data interface{}) *Node {
	return &Node{data: data}
}

func (p *Node) Next() *Node {
	if p.next == nil || p.next.data == nil {
		return nil
	}
	return p.next
}

func (p *Node) Prev() *Node {
	if p.prev == nil || p.prev.data == nil {
		return nil
	}
	return p.prev
}

func (p *Node) Data() interface{} {
	return p.data
}
