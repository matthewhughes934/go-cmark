package gfm

/*
#include "cmark-gfm.h"
*/
import "C"

type Node struct {
	node *C.struct_cmark_node
}
