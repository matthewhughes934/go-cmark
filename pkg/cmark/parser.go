package cmark

/*
#include "cmark.h"
#include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"unsafe"
)

type Parser struct {
	parser *C.cmark_parser
}

type ParserOpt int

const (
	// Default options.
	ParserOptDefault ParserOpt = C.CMARK_OPT_DEFAULT

	// Include a `data-sourcepos` attribute on all block elements.
	ParserOptSourcePos ParserOpt = C.CMARK_OPT_SOURCEPOS

	// Render `softbreak` elements as hard line breaks.
	ParserOptHardBreaks ParserOpt = C.CMARK_OPT_HARDBREAKS

	//  Render raw HTML and unsafe links (`javascript:`, `vbscript:`,
	// `file:`, and `data:`, except for `image/png`, `image/gif`,
	// `image/jpeg`, or `image/webp` mime types).  By default,
	// raw HTML is replaced by a placeholder HTML comment. Unsafe
	// links are replaced by empty strings.
	ParserOptUnsafe ParserOpt = C.CMARK_OPT_UNSAFE

	// Render `softbreak` elements as spaces.
	ParserOptNoBreaks ParserOpt = C.CMARK_OPT_NOBREAKS

	// Legacy option (no effect).
	ParserOptNormalize ParserOpt = C.CMARK_OPT_NORMALIZE

	//  Validate UTF-8 in the input before parsing, replacing illegal
	// sequences with the replacement character U+FFFD.
	ParserOptValidateUTF8 ParserOpt = C.CMARK_OPT_VALIDATE_UTF8

	// Convert straight quotes to curly, --- to em dashes, -- to en dashes.
	ParserOptSmart ParserOpt = C.CMARK_OPT_SMART
)


// free wraps cmark_parser_free
func (parser *Parser) free() { //go-cov:skip
	C.cmark_parser_free(parser.parser)
}

// NewParser wraps cmark_parser_new
// Creates a new parser object.
func NewParser(options ParserOpt) *Parser {
	parser := &Parser{parser: C.cmark_parser_new(C.int(options))}
	runtime.SetFinalizer(parser, (*Parser).free)

	return parser
}

// Feed wraps cmark_parser_feed
// Feeds a string from 'text' to 'parser'.
func (parser *Parser) Feed(text string) {
	Cstr := C.CString(text)
	defer C.free(unsafe.Pointer(Cstr))

	C.cmark_parser_feed(parser.parser, Cstr, C.size_t(len(text)))
}

// Finish wraps cmark_parser_finish
// Finish parsing and return a pointer to a tree of nodes.
func (parser *Parser) Finish() *Node {
	return &Node{
		node: C.cmark_parser_finish(parser.parser),
	}
}

// ParseDocument wraps cmark_parse_document
// Parse a CommonMark document in 'document' and returns a pointer to a tree of nodes.
// The returned [cmark.Node] has a finalizer set that will call
// `cmark_node_free` which will free the memory allocated for the node and any
// of its children
func ParseDocument(document string, options ParserOpt) *Node {
	str := C.CString(document)
	defer C.free(unsafe.Pointer(str))

	node := &Node{
		node: C.cmark_parse_document(str, C.size_t(len(document)), C.int(options)),
	}
	runtime.SetFinalizer(node, (*Node).free)

	return node
}
