package mecab

// #include <mecab.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
)

type Node struct {
	head    *C.mecab_node_t
	current *C.struct_mecab_node_t
}

var StopIteration = errors.New("StopIteration")

func (node *Node) Next() error {
	node.current = node.current.next
	if node.current == nil {
		return StopIteration
	}
	return nil
}

func (node *Node) Surface() string {
	current := node.current
	return C.GoString(current.surface)[:int(current.length)]
}

func (node *Node) Feature() string {
	return C.GoString(node.current.feature)
}
