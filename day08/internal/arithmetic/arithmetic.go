package arithmetic

import (
	"github.com/pkg/errors"
	"strconv"
	"unsafe"
)

func GetElement(arr []int, idx int) (int, error) {
	if idx < 0 {
		return 0, errors.New("index is negative")
	} else if idx >= len(arr) {
		return 0, errors.New("index " + strconv.Itoa(idx) + " is out of range")
	} else if len(arr) == 0 {
		return 0, errors.New("incorrect array")
	}

	ptrStart := unsafe.Pointer(&arr[0])
	itemSize := unsafe.Sizeof(arr[0])
	element := *(*int)(unsafe.Add(ptrStart, uintptr(idx)*itemSize))
	return element, nil
}
