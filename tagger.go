package mecab

// #include <mecab.h>
// #include <stdlib.h>
import "C"

type Tagger struct {
	tagger *C.mecab_t
}

func (tg *Tagger) Destroy() {
	C.mecab_destroy(tg.tagger)
}

func (tg *Tagger) Parse(lt *Lattice) string {
	C.mecab_parse_lattice(tg.tagger, lt.lattice)
	return C.GoString(C.mecab_lattice_tostr(lt.lattice))
}

func (tg *Tagger) ParseToNode(lt *Lattice) *Node {
	C.mecab_parse_lattice(tg.tagger, lt.lattice)
	node := C.mecab_lattice_get_bos_node(lt.lattice)
	return &Node{node, node}
}
