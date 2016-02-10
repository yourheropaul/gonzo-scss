package libsass

/*
#include "sass_context.h"
*/
import "C"

type Context interface {
	Options() Options
	OutputString() string
	Error() error
	SourceMap() string
	IncludedFiles() []string

	//TODO: XXX???
	//TakeError()?
}

var _ Context = &context{}

type context C.struct_Sass_Context

func (c *context) Options() Options {
	cc := (*C.struct_Sass_Context)(c)
	return (*options)(C.sass_context_get_options(cc))
}

func (c *context) OutputString() string {
	cc := (*C.struct_Sass_Context)(c)
	return C.GoString(C.sass_context_get_output_string(cc))
}

func (c *context) Error() error {
	cc := (*C.struct_Sass_Context)(c)
	return newError(cc)
}

func (c *context) SourceMap() string {
	cc := (*C.struct_Sass_Context)(c)
	return C.GoString(C.sass_context_get_source_map_string(cc))
}

func (c *context) IncludedFiles() []string {
	//hdr := reflect.SliceHeader{
	//	Data: uintptr(unsafe.Pointer(files)),
	//	Len:  length,
	//	Cap:  length,
	//}

	return []string{}
	//*(*[]string)(unsafe.Pointer(&hdr))

}
