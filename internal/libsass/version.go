package libsass

/*
  //Pending golang/go/issues/12281
# //cgo CPPFLAGS: -D LIBSASS_VERSION="3.2.4-0ae11a4"
#include "sass_context.h"
*/
import "C"

func Version() string {
	return C.GoString(C.libsass_version())
}
