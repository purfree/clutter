package singly

type Node struct {
	data interface{}
	next *Node
}

func newNode(data interface{}) *Node {
	return &Node{data: data}
}

func (p *Node) Next() *Node {
	return p.next
}

func (p *Node) Data() interface{} {
	return p.data
}
