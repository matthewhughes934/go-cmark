// Code generated by ../../scripts/copy-lib ../../cmark-gfm-src at 1e230827a584ebc9938c3eadc5059c55ef3c9abf DO NOT EDIT.
// See ../../cmark-gfm-src/COPYING for license
#ifndef CMARK_GFM_TABLE_H
#define CMARK_GFM_TABLE_H

#include "cmark-gfm-core-extensions.h"


extern cmark_node_type CMARK_NODE_TABLE, CMARK_NODE_TABLE_ROW,
    CMARK_NODE_TABLE_CELL;

cmark_syntax_extension *create_table_extension(void);

#endif
