package libsass

/*
#include "sass_context.h"
*/
import "C"
import "errors"

type Compiler interface {
	Parse() error
	Execute() error
	Destroy() error
}

type compiler C.struct_Sass_Compiler

func (c *compiler) Parse() error {
	errc := C.sass_compiler_parse((*C.struct_Sass_Compiler)(c))

	//TODO: Proper error handling.
	if errc != 0 {
		return errors.New("Something went wrong.")
	}
	return nil
}

func (c *compiler) Execute() error {
	errc := C.sass_compiler_execute((*C.struct_Sass_Compiler)(c))
	//TODO: Proper error handling.
	if errc != 0 {
		return errors.New("Something went wrong.")
	}
	return nil
}

func (c *compiler) Destroy() error {
	C.sass_delete_compiler((*C.struct_Sass_Compiler)(c))
	//XXX
	return nil
}
