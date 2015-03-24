package resize

/*
#include <stdlib.h>
#cgo linux  pkg-config: opencv
#cgo darwin pkg-config: opencv
#cgo freebsd pkg-config: opencv
#cgo windows LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
int resize(char* fp, char* sp, int w, int h);
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

/*
Args
 fp:the target image file path.
 sp:the save path for resized image.
 w:target width.
 h:target height.
Return
 0:success
 -1:the target file not found.
 -2:can save the file to sp.
*/
func Resize(fp, sp string, w, h int) error {
	fp_, sp_ := C.CString(fp), C.CString(sp)
	defer func() {
		C.free(unsafe.Pointer(fp_))
		C.free(unsafe.Pointer(sp_))
	}()
	ret := int(C.resize(fp_, sp_, C.int(w), C.int(h)))
	switch ret {
	case -1:
		return errors.New(fmt.Sprintf("read file %v fail", fp))
	case -2:
		return errors.New(fmt.Sprintf("write file %v fail", sp))
	default:
		return nil
	}
}
