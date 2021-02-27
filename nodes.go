package okdoc

import "fmt"

type Content struct {
	Nodes []*Node
}

type Node struct {
	Type string
	Raw  string
	Line uint
}

func (n *Node) String() string {
	return fmt.Sprintf("L%d %s(%q)", n.Line, n.Type, n.Raw)
}
