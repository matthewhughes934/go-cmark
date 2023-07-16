// Package gfm provides Go bindings for the cmark-gfm library
package gfm

//go:generate ../../scripts/copy-lib ../../cmark-gfm-src

/*
#include "cmark-gfm.h"
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

// RenderCommonMark wraps `cmark_render_commonmark`.
// Returns the tree under `root` rendered as a commonmark document
func RenderCommonMark(root *Node, options ParserOpt, width int) string {
	cStr := C.cmark_render_commonmark(root.node, C.int(options), C.int(width))
	defer C.free(unsafe.Pointer(cStr))
	return C.GoString(cStr)
}
