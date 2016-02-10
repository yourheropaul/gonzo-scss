// +build !shared

package libsass

/*
#cgo CPPFLAGS: -D STATIC=1
#cgo CXXFLAGS: -g -std=c++11 -I ./libsass/include
#cgo CFLAGS:  -Ilibsass/include
#cgo LDFLAGS: -lstdc++ -lm
#cgo darwin linux LDFLAGS: -ldl
*/
import "C"
