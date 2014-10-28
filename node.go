package mecab

// #include <mecab.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
)

type Node struct {
	// pointer to the first node
	head *C.mecab_node_t
	// pointer to he next node
	current *C.struct_mecab_node_t
}

var StopIteration = errors.New("StopIteration")

// proceed to the next node.
// if current node is last, this method returns StopIteration error.
func (node *Node) Next() error {
	node.current = node.current.next
	if node.current == nil {
		return StopIteration
	}
	return nil
}

// surface string
func (node *Node) Surface() string {
	current := node.current
	return C.GoString(current.surface)[:int(current.length)]
}

// feature string
func (node *Node) Feature() string {
	return C.GoString(node.current.feature)
}

// unique node id
func (node *Node) Id() int {
	return int(node.current.id)
}

// length of the surface form.
func (node *Node) Length() int {
	return int(node.current.length)
}

// length of the surface form including white space before the morph.
func (node *Node) Rlength() int {
	return int(node.current.rlength)
}

// right attribute id
func (node *Node) RcAttr() int {
	return int(node.current.rcAttr)
}

// left attribute id
func (node *Node) LcAttr() int {
	return int(node.current.lcAttr)
}

// unique part of speech id.
func (node *Node) Posid() int {
	return int(node.current.posid)
}

// character type
func (node *Node) Char_type() int {
	return int(node.current.char_type)
}

// status of this model.
func (node *Node) Stat() int {
	return int(node.current.stat)
}

// set 1 if this node is best node.
func (node *Node) Isbest() int {
	return int(node.current.isbest)
}

// forward accumulative log summation.
func (node *Node) Alpha() float32 {
	return float32(node.current.alpha)
}

// backward accumulative log summation.
func (node *Node) Beta() float32 {
	return float32(node.current.beta)
}

// marginal probability.
func (node *Node) Prob() float32 {
	return float32(node.current.prob)
}

// word cost.
func (node *Node) Wcost() int {
	return int(node.current.wcost)
}

// best accumulative cost from bos node to this node.
func (node *Node) Cost() int {
	return int(node.current.cost)
}
