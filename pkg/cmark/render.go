package cmark

/*
#include "cmark.h"
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

// RenderOpts provides options for configuring rendering
type RenderOpts struct {
	o C.int
}

func NewRenderOpts() *RenderOpts {
	return &RenderOpts{o: C.CMARK_OPT_DEFAULT}
}

// WithSourcePos maps to C.CMARK_OPT_SOURCEPOS.
// Includes a `data-sourcepos` attribute on all block elements.
func (o *RenderOpts) WithSourcePos() *RenderOpts {
	o.o |= C.CMARK_OPT_SOURCEPOS
	return o
}

// WithHarbreaks maps to C.CMARK_OPT_HARDBREAKS.
// Renders `softbreak` elements as hard line breaks.
func (o *RenderOpts) WithHarbreaks() *RenderOpts {
	o.o |= C.CMARK_OPT_HARDBREAKS
	return o
}

// WithUnsafe maps to C.CMARK_OPT_UNSAFE.
// Renders raw HTML and unsafe links (`javascript:`, `vbscript:`,
// `file:`, and `data:`, except for `image/png`, `image/gif`,
// `image/jpeg`, or `image/webp` mime types).  By default,
// raw HTML is replaced by a placeholder HTML comment. Unsafe
// links are replaced by empty strings.
func (o *RenderOpts) WithUnsafe() *RenderOpts {
	o.o |= C.CMARK_OPT_UNSAFE
	return o
}

// WithNoBreaks maps to C.CMARK_OPT_NOBREAKS.
// Renders `softbreak` elements as spaces.
func (o *RenderOpts) WithNoBreaks() *RenderOpts {
	o.o |= C.CMARK_OPT_NOBREAKS
	return o
}

// RenderCommonMark wraps `cmark_render_commonmark`.
// Returns the tree under `root` rendered as a commonmark document
func RenderCommonMark(root *Node, opts *RenderOpts, width int) string {
	cStr := C.cmark_render_commonmark(root.node, opts.o, C.int(width))
	defer C.free(unsafe.Pointer(cStr))
	return C.GoString(cStr)
}

// RenderHTML wraps cmark_render_html.
// Renders the tree under 'root' as an HTML fragment.
// It is up to the user to add an appropriate header and footer.
func RenderHTML(root *Node, opts *RenderOpts) string {
	cStr := C.cmark_render_html(root.node, opts.o)
	defer C.free(unsafe.Pointer(cStr))
	return C.GoString(cStr)
}
