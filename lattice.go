package mecab

// #include <mecab.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

type Lattice struct {
	lattice  *C.mecab_lattice_t
	sentence *C.char
}

func (lt *Lattice) Destroy() {
	C.mecab_lattice_destroy(lt.lattice)
	C.free(unsafe.Pointer(lt.sentence))
}
