package solution

import (
	"errors"
)

var (
	ErrNoValidChild = errors.New("no valid child available")
)

type Status int

const (
	NEW Status = iota
	FULL
	DEAD
	HIDE
)

// Node represents a node in the search tree
type Node struct {
	Root     bool
	parent   *Node
	Value    int
	Status   Status
	Win      int
	Loss     int
	Depth    int
	ObjFunc  float64
	Children []*Node
}

// CreateRoot returns a node with default root node settings
func CreateRoot(v int) *Node {
	return &Node{
		Root:   true,
		Value:  v,
		Depth:  0,
		Status: NEW,
	}
}

// CreateDeepRoot returns a node with default root node settings and a tree above it
func CreateDeepRoot(vals []int) *Node {
	n := &Node{
		Root:   false,
		Value:  vals[0],
		Depth:  0,
		Status: NEW,
	}
	for _, v := range vals[1:] {
		c := n.CreateChild(v)
		n = c
	}
	n.Root = true
	return n
}

// CreateChild adds a node as the child with the value v, returns the
// new child
func (n *Node) CreateChild(v int) *Node {
	n.Children = append(n.Children, &Node{
		parent: n,
		Value:  v,
		Depth:  n.Depth + 1,
		Status: NEW,
	})
	return n.LastChild()
}

// GenerateChild attempts to create a child where the childs value is determined by
// function f, if no child can be generated returns ErrNoValidChild
func (n *Node) GenerateChild(clubs int, f func(n *Node, clubs int) (int, float64, error)) (*Node, error) {
	v, o, err := f(n, clubs)
	if err != nil {
		return nil, err
	}
	x := n.CreateChild(v)
	x.ObjFunc = o
	return x, nil
}

// RootNode returns the root node of the branch this node is on
func (n *Node) RootNode() *Node {
	if n.Root {
		return n
	}
	p := n
	for {
		if p.Root {
			return p
		}
		p = p.Parent()
	}
}

// Parent returns the parent node if this is not a root node. Otherwise
// returns nil
func (n *Node) Parent() *Node {
	return n.parent
}

// LastChild returns the youngest child or nil if there are no children
func (n *Node) LastChild() *Node {
	l := len(n.Children)
	if l == 0 {
		return nil
	}
	return n.Children[l-1]
}

func (n *Node) RecurseWin() {
	n.Win += 1
	p := n // Avoid mutating this node
	for {
		p = p.Parent()
		if p == nil {
			break
		}
		p.Win += 1
	}
}

func (n *Node) RecurseLoss() {
	n.Loss += 1
	p := n // Avoid mutating this node
	for {
		p = p.Parent()
		if p == nil {
			break
		}
		p.Loss += 1
	}
}

// Tries total number of wins and losses on this branch which is actually
// the wins, losses of the root node
func (n *Node) Tries() int {
	if n.Root {
		return n.Win + n.Loss
	}
	return n.RootNode().Tries()
}

// Kill will mark this node as dead and remove the pointer to its children to le the GC clear memory
func (n *Node) Kill() {
	n.Status = DEAD
	n.Children = nil
}

// Value returns a slice of the values obtained by travelling up the
// tree to the root
func (n *Node) Values(f func(int)) {
	f(n.Value)
	p := n // Avoid mutating this node
	for {
		p = p.Parent()
		if p == nil {
			break
		}
		f(p.Value)
	}
}

// ReverseValues returns a slice of Values in reverse order
func (n *Node) ReverseValues() []int {
	vals := make([]int, 0, n.Depth+1)
	n.Values(func(v int) {
		vals = append(vals, v)
	})
	rVals := make([]int, len(vals))
	for i, v := range vals {
		rVals[len(vals)-i-1] = v
	}
	return rVals
}

// Branch returns a slice of all nodes in this branch of the tree from
// the child to the root
func (n *Node) Branch() []*Node {
	nodes := make([]*Node, 0, n.Depth+1)
	nodes = append(nodes, n)
	p := n // Avoid mutating this node
	for {
		p = p.Parent()
		if p == nil {
			break
		}
		nodes = append(nodes, p)
	}
	return nodes
}
