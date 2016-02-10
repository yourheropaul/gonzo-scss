package libsass

/*
#include "sass_context.h"
*/
import "C"
import "fmt"

type Error interface {
	Status() error
	JSON() string
	Error() string //Message
	Text() string
	File() string
	Line() int64
	Column() int64
}

func newError(c *C.struct_Sass_Context) Error {
	return &err{ctx: c}
}

//TODO: Should we copy all the errors to avoid calls on
// Sass_Context after Close()? Probably.
type err struct {
	ctx *C.struct_Sass_Context
}

func (e *err) Status() error {
	errc := C.sass_context_get_error_status(e.ctx)
	//TODO: Proper error handling.
	if errc != 0 {
		return fmt.Errorf("Something went wrong: Code %s", errc)
	}
	return nil
}

func (e *err) JSON() string {
	return C.GoString(C.sass_context_get_error_json(e.ctx))
}

func (e *err) Text() string {
	return C.GoString(C.sass_context_get_error_text(e.ctx))
}

func (e *err) Error() string {
	return C.GoString(C.sass_context_get_error_message(e.ctx))
}

func (e *err) File() string {
	return C.GoString(C.sass_context_get_error_file(e.ctx))
}

func (e *err) Line() int64 {
	return int64(C.sass_context_get_error_line(e.ctx))
}
func (e *err) Column() int64 {
	return int64(C.sass_context_get_error_column(e.ctx))
}
