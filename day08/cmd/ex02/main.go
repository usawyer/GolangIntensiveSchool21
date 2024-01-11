package main

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "window.h"
// #include "application.h"
import "C"
import (
	"unsafe"
)

func main() {
	C.InitApplication()
	tittle := C.CString("School 21")
	defer C.free(unsafe.Pointer(tittle))
	windowPtr := C.Window_Create(0, 0, 300, 200, tittle)
	defer C.free(unsafe.Pointer(windowPtr))
	C.Window_MakeKeyAndOrderFront(windowPtr)
	C.RunApplication()
}
