package cursor

import (
	"syscall"
	"unsafe"
)

type POINT struct {
	X, Y int32
}

var (
	user32       = syscall.NewLazyDLL("user32.dll")
	getCursorPos = user32.NewProc("GetCursorPos")
)

// Getcoords Func Offset Cords
func Getcoords(point *POINT) bool {

	getCursorPos.Call(uintptr(unsafe.Pointer(point)), 0, 0, 0)
	return true
}

/*ret, _, _ := syscall.Syscall(
getCursorPos.Addr(), 1, uintptr(unsafe.Pointer(point)), 0, 0)*/
