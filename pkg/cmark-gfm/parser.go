package gfm

/*
#include <stdlib.h>
#include "cmark-gfm.h"
#include "cmark-gfm-extension_api.h"
#include "parser.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

type Parser struct {
	parser *C.cmark_parser
}

// WithValidateUTF8 maps to C.CMARK_OPT_VALIDATE_UTF8.
// Validates UTF-8 in the input before parsing, replacing illegal
// sequences with the replacement character U+FFFD.
func (p *Parser) WithValidateUTF8() *Parser {
	p.parser.options |= C.CMARK_OPT_VALIDATE_UTF8
	return p
}

// WithSmart maps to C.CMARK_OPT_SMART.
// Converts straight quotes to curly, --- to em dashes, -- to en dashes.
func (p *Parser) WithSmart() *Parser {
	p.parser.options |= C.CMARK_OPT_SMART
	return p
}

// WithFoonotes maps to c.CMARK_OPT_FOOTNOTES
// Parses footnotes
func (p *Parser) WithFoonotes() *Parser {
	p.parser.options |= C.CMARK_OPT_FOOTNOTES
	return p
}

// free wraps cmark_parser_free
func (parser *Parser) free() { //go-cov:skip
	C.cmark_parser_free(parser.parser)
}

// NewParser wraps cmark_parser_new.
// Creates a new parser object.
func NewParser() *Parser {
	parser := &Parser{parser: C.cmark_parser_new(C.CMARK_OPT_DEFAULT)}
	runtime.SetFinalizer(parser, (*Parser).free)

	return parser
}

// Feed wraps cmark_parser_feed.
// Feeds a string from 'text' to 'parser'.
func (parser *Parser) Feed(text string) {
	Cstr := C.CString(text)
	defer C.free(unsafe.Pointer(Cstr))

	C.cmark_parser_feed(parser.parser, Cstr, C.size_t(len(text)))
}

// Finish wraps cmark_parser_finish.
// Finish parsing and return a pointer to a tree of nodes.
func (parser *Parser) Finish() *Node {
	return &Node{
		node: C.cmark_parser_finish(parser.parser),
	}
}

// ParseDocument parses a CommonMark document in 'document' and returns a pointer to a tree of nodes.
// The returned [cmark.Node] has a finalizer set that will call
// `cmark_node_free` which will free the memory allocated for the node and any
// of its children
func (parser *Parser) ParseDocument(document string) *Node {
	parser.Feed(document)
	node := parser.Finish()
	runtime.SetFinalizer(node, (*Node).free)

	return node
}

// AttachSyntaxExtension wraps cmark_parser_attach_syntax_extension.
// Attaches the given [SyntaxExtension] to the parser.
func (parser *Parser) AttachSyntaxExtension(extension *SyntaxExtension) {
	C.cmark_parser_attach_syntax_extension(parser.parser, extension.ext)
}

// SyntaxExtensions accesses cmark_parser.syntax_extensions.
// Returns the [SyntaxExtensionList] attached to the parser.
func (parser *Parser) SyntaxExtensions() *SyntaxExtensionList {
	return &SyntaxExtensionList{llist: parser.parser.syntax_extensions}
}
