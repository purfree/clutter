package circular

type Node struct {
	data interface{}
	next *Node
	prev *Node
	list *List
}

func newNode(data interface{}) *Node {
	return &Node{data: data}
}

func (p *Node) Next() *Node {
	return p.next
}

func (p *Node) Prev() *Node {
	return p.prev
}

func (p *Node) Data() interface{} {
	return p.data
}
