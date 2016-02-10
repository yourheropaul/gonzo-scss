package libsass

/*
#include "sass_context.h"
*/
import "C"

type DataContext interface {
	Context() Context
	SetOptions(opt Options)
	Compiler() Compiler

	Destroy() error
}

func NewDataContext(data string) DataContext {
	return (*datacontext)(C.sass_make_data_context(C.CString(data)))
}

type datacontext C.struct_Sass_Data_Context

func (dc *datacontext) Context() Context {
	cc := (*C.struct_Sass_Data_Context)(dc)
	return (*context)(C.sass_data_context_get_context(cc))
}

func (dc *datacontext) SetOptions(opt Options) {
	C.sass_data_context_set_options(
		(*C.struct_Sass_Data_Context)(dc),
		opt.optc(),
	)
}

func (dc *datacontext) Compiler() Compiler {
	cc := (*C.struct_Sass_Data_Context)(dc)
	return (*compiler)(C.sass_make_data_compiler(cc))
}

func (dc *datacontext) Destroy() error {
	cc := (*C.struct_Sass_Data_Context)(dc)
	C.sass_delete_data_context(cc)
	//XXX: ?
	return nil
}
