package mecab

// #include <mecab.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type MeCab struct {
	model *C.mecab_model_t
}

func New(option ...string) (*MeCab, error) {
	var opt string
	if len(option) != 0 {
		opt = option[0]
	}
	_opt := C.CString(opt)
	defer C.free(unsafe.Pointer(_opt))
	_model := C.mecab_model_new2(_opt)
	if _model == nil {
		return nil, fmt.Errorf("mecab_model is not created: %v", C.GoString(C.mecab_strerror(nil)))
	}
	return &MeCab{
		model: _model,
	}, nil
}

func (m *MeCab) Destroy() {
	C.mecab_model_destroy(m.model)
}

func (m *MeCab) NewLattice(input string) (*Lattice, error) {
	_lt := C.mecab_model_new_lattice(m.model)
	if _lt == nil {
		return nil, errors.New("mecab_lattice is not created.")
	}
	_input := C.CString(input)
	C.mecab_lattice_set_sentence(_lt, _input)
	return &Lattice{
		lattice:  _lt,
		sentence: _input,
	}, nil
}

func (m *MeCab) NewTagger() (*Tagger, error) {
	_tg := C.mecab_model_new_tagger(m.model)
	if _tg == nil {
		return nil, errors.New("mecab_tagger is not created.")
	}
	return &Tagger{
		tagger: _tg,
	}, nil
}
