package gfm

/*
#include <stdlib.h>
#include "cmark-gfm.h"
#include "plugin.h"
#include "registry.h"
#include "cmark-gfm-extension_api.h"
#include "cmark-gfm-core-extensions.h"
*/
import "C"

import (
	"unsafe"
)

// CoreExtensionsEnsureRegistered wraps
// cmark_gfm_core_extensions_ensure_registered.
// Checks if core extensions have been registered, and registers them if not.
//
// You should call [ReleasePlugins] to free the memory allocated. Usage:
//
//	CoreExtensionsEnsureRegistered()
//	defer ReleasePlugins()
func CoreExtensionsEnsureRegistered() {
	C.cmark_gfm_core_extensions_ensure_registered()
}

// ReleasePlugins wraps cmark_release_plugins.
// Frees the memory allocated when registering psyntax extensions.
func ReleasePlugins() { //go-cov:skip this is only run at the end of tests
	C.cmark_release_plugins()
}

type SyntaxExtension struct {
	ext *C.cmark_syntax_extension
}

type SyntaxExtensionList struct {
	llist *C.cmark_llist
}

// FindSyntaxExtension wraps cmark_find_syntax_extension.
// Finds the syntax extension with the given name, or 'nil' if not such extension is registered
func FindSyntaxExtension(name string) *SyntaxExtension {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	ext := C.cmark_find_syntax_extension(cs)
	if ext == nil {
		return nil
	}
	return &SyntaxExtension{ext: ext}
}
