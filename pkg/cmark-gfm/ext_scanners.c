// Code generated by ../../scripts/copy-lib ../../cmark-gfm-src at 587a12bb54d95ac37241377e6ddc93ea0e45439b DO NOT EDIT.
// See ../../cmark-gfm-src/COPYING for license
/* Generated by re2c 1.3 */

#include "ext_scanners.h"
#include <stdlib.h>

bufsize_t _ext_scan_at(bufsize_t (*scanner)(const unsigned char *),
                       unsigned char *ptr, int len, bufsize_t offset) {
  bufsize_t res;

  if (ptr == NULL || offset >= len) {
    return 0;
  } else {
    unsigned char lim = ptr[len];

    ptr[len] = '\0';
    res = scanner(ptr + offset);
    ptr[len] = lim;
  }

  return res;
}

bufsize_t _scan_table_start(const unsigned char *p) {
  const unsigned char *marker = NULL;
  const unsigned char *start = p;

  {
    unsigned char yych;
    static const unsigned char yybm[] = {
        0, 0,   0, 0, 0, 0, 0, 0, 0, 64, 0,  64, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0,   0, 0, 0, 0, 0, 0, 0, 0,  64, 0,  0,  0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 128, 0, 0, 0, 0, 0, 0, 0, 0,  0,  0,  0,  0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0,   0, 0, 0, 0, 0, 0, 0, 0,  0,  0,  0,  0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0,   0, 0, 0, 0, 0, 0, 0, 0,  0,  0,  0,  0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0,   0, 0, 0, 0, 0, 0, 0, 0,  0,  0,  0,  0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0,   0, 0, 0, 0, 0, 0, 0, 0,  0,  0,  0,  0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0,   0, 0, 0, 0, 0, 0, 0, 0,  0,  0,  0,  0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0,   0, 0, 0, 0, 0, 0, 0, 0,  0,  0,  0,  0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0,   0, 0, 0, 0, 0, 0, 0, 0,  0,  0,  0,  0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0,   0, 0, 0, 0, 0, 0, 0, 0,  0,  0,  0,  0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0,   0, 0, 0, 0, 0, 0, 0, 0,  0,  0,  0,  0,
    };
    yych = *p;
    if (yych <= ' ') {
      if (yych <= '\n') {
        if (yych == '\t')
          goto yy4;
      } else {
        if (yych <= '\f')
          goto yy4;
        if (yych >= ' ')
          goto yy4;
      }
    } else {
      if (yych <= '9') {
        if (yych == '-')
          goto yy5;
      } else {
        if (yych <= ':')
          goto yy6;
        if (yych == '|')
          goto yy4;
      }
    }
    ++p;
  yy3 : { return 0; }
  yy4:
    yych = *(marker = ++p);
    if (yybm[0 + yych] & 64) {
      goto yy7;
    }
    if (yych == '-')
      goto yy10;
    if (yych == ':')
      goto yy12;
    goto yy3;
  yy5:
    yych = *(marker = ++p);
    if (yybm[0 + yych] & 128) {
      goto yy10;
    }
    if (yych <= ' ') {
      if (yych <= 0x08)
        goto yy3;
      if (yych <= '\r')
        goto yy14;
      if (yych <= 0x1F)
        goto yy3;
      goto yy14;
    } else {
      if (yych <= ':') {
        if (yych <= '9')
          goto yy3;
        goto yy13;
      } else {
        if (yych == '|')
          goto yy14;
        goto yy3;
      }
    }
  yy6:
    yych = *(marker = ++p);
    if (yybm[0 + yych] & 128) {
      goto yy10;
    }
    goto yy3;
  yy7:
    yych = *++p;
    if (yybm[0 + yych] & 64) {
      goto yy7;
    }
    if (yych == '-')
      goto yy10;
    if (yych == ':')
      goto yy12;
  yy9:
    p = marker;
    goto yy3;
  yy10:
    yych = *++p;
    if (yybm[0 + yych] & 128) {
      goto yy10;
    }
    if (yych <= 0x1F) {
      if (yych <= '\n') {
        if (yych <= 0x08)
          goto yy9;
        if (yych <= '\t')
          goto yy13;
        goto yy15;
      } else {
        if (yych <= '\f')
          goto yy13;
        if (yych <= '\r')
          goto yy17;
        goto yy9;
      }
    } else {
      if (yych <= ':') {
        if (yych <= ' ')
          goto yy13;
        if (yych <= '9')
          goto yy9;
        goto yy13;
      } else {
        if (yych == '|')
          goto yy18;
        goto yy9;
      }
    }
  yy12:
    yych = *++p;
    if (yybm[0 + yych] & 128) {
      goto yy10;
    }
    goto yy9;
  yy13:
    yych = *++p;
  yy14:
    if (yych <= '\r') {
      if (yych <= '\t') {
        if (yych <= 0x08)
          goto yy9;
        goto yy13;
      } else {
        if (yych <= '\n')
          goto yy15;
        if (yych <= '\f')
          goto yy13;
        goto yy17;
      }
    } else {
      if (yych <= ' ') {
        if (yych <= 0x1F)
          goto yy9;
        goto yy13;
      } else {
        if (yych == '|')
          goto yy18;
        goto yy9;
      }
    }
  yy15:
    ++p;
    { return (bufsize_t)(p - start); }
  yy17:
    yych = *++p;
    if (yych == '\n')
      goto yy15;
    goto yy9;
  yy18:
    yych = *++p;
    if (yybm[0 + yych] & 128) {
      goto yy10;
    }
    if (yych <= '\r') {
      if (yych <= '\t') {
        if (yych <= 0x08)
          goto yy9;
        goto yy18;
      } else {
        if (yych <= '\n')
          goto yy15;
        if (yych <= '\f')
          goto yy18;
        goto yy17;
      }
    } else {
      if (yych <= ' ') {
        if (yych <= 0x1F)
          goto yy9;
        goto yy18;
      } else {
        if (yych == ':')
          goto yy12;
        goto yy9;
      }
    }
  }
}

bufsize_t _scan_table_cell(const unsigned char *p) {
  const unsigned char *marker = NULL;
  const unsigned char *start = p;

  {
    unsigned char yych;
    unsigned int yyaccept = 0;
    static const unsigned char yybm[] = {
        64, 64, 64,  64, 64, 64, 64, 64, 64, 64, 0,  64, 64, 0,  64, 64, 64, 64,
        64, 64, 64,  64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
        64, 64, 64,  64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
        64, 64, 64,  64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
        64, 64, 64,  64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
        64, 64, 128, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
        64, 64, 64,  64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 0,  64,
        64, 64, 0,   0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,
        0,  0,  0,   0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,
        0,  0,  0,   0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,
        0,  0,  0,   0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,
        0,  0,  0,   0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,
        0,  0,  0,   0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,
        0,  0,  0,   0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,
        0,  0,  0,   0,
    };
    yych = *p;
    if (yybm[0 + yych] & 64) {
      goto yy22;
    }
    if (yych <= 0xEC) {
      if (yych <= 0xC1) {
        if (yych <= '\r')
          goto yy25;
        if (yych <= '\\')
          goto yy27;
        goto yy25;
      } else {
        if (yych <= 0xDF)
          goto yy29;
        if (yych <= 0xE0)
          goto yy30;
        goto yy31;
      }
    } else {
      if (yych <= 0xF0) {
        if (yych <= 0xED)
          goto yy32;
        if (yych <= 0xEF)
          goto yy31;
        goto yy33;
      } else {
        if (yych <= 0xF3)
          goto yy34;
        if (yych <= 0xF4)
          goto yy35;
        goto yy25;
      }
    }
  yy22:
    yyaccept = 0;
    yych = *(marker = ++p);
    if (yybm[0 + yych] & 64) {
      goto yy22;
    }
    if (yych <= 0xEC) {
      if (yych <= 0xC1) {
        if (yych <= '\r')
          goto yy24;
        if (yych <= '\\')
          goto yy27;
      } else {
        if (yych <= 0xDF)
          goto yy36;
        if (yych <= 0xE0)
          goto yy38;
        goto yy39;
      }
    } else {
      if (yych <= 0xF0) {
        if (yych <= 0xED)
          goto yy40;
        if (yych <= 0xEF)
          goto yy39;
        goto yy41;
      } else {
        if (yych <= 0xF3)
          goto yy42;
        if (yych <= 0xF4)
          goto yy43;
      }
    }
  yy24 : { return (bufsize_t)(p - start); }
  yy25:
    ++p;
  yy26 : { return 0; }
  yy27:
    yyaccept = 0;
    yych = *(marker = ++p);
    if (yybm[0 + yych] & 128) {
      goto yy27;
    }
    if (yych <= 0xDF) {
      if (yych <= '\f') {
        if (yych == '\n')
          goto yy24;
        goto yy22;
      } else {
        if (yych <= '\r')
          goto yy24;
        if (yych <= 0x7F)
          goto yy22;
        if (yych <= 0xC1)
          goto yy24;
        goto yy36;
      }
    } else {
      if (yych <= 0xEF) {
        if (yych <= 0xE0)
          goto yy38;
        if (yych == 0xED)
          goto yy40;
        goto yy39;
      } else {
        if (yych <= 0xF0)
          goto yy41;
        if (yych <= 0xF3)
          goto yy42;
        if (yych <= 0xF4)
          goto yy43;
        goto yy24;
      }
    }
  yy29:
    yych = *++p;
    if (yych <= 0x7F)
      goto yy26;
    if (yych <= 0xBF)
      goto yy22;
    goto yy26;
  yy30:
    yyaccept = 1;
    yych = *(marker = ++p);
    if (yych <= 0x9F)
      goto yy26;
    if (yych <= 0xBF)
      goto yy36;
    goto yy26;
  yy31:
    yyaccept = 1;
    yych = *(marker = ++p);
    if (yych <= 0x7F)
      goto yy26;
    if (yych <= 0xBF)
      goto yy36;
    goto yy26;
  yy32:
    yyaccept = 1;
    yych = *(marker = ++p);
    if (yych <= 0x7F)
      goto yy26;
    if (yych <= 0x9F)
      goto yy36;
    goto yy26;
  yy33:
    yyaccept = 1;
    yych = *(marker = ++p);
    if (yych <= 0x8F)
      goto yy26;
    if (yych <= 0xBF)
      goto yy39;
    goto yy26;
  yy34:
    yyaccept = 1;
    yych = *(marker = ++p);
    if (yych <= 0x7F)
      goto yy26;
    if (yych <= 0xBF)
      goto yy39;
    goto yy26;
  yy35:
    yyaccept = 1;
    yych = *(marker = ++p);
    if (yych <= 0x7F)
      goto yy26;
    if (yych <= 0x8F)
      goto yy39;
    goto yy26;
  yy36:
    yych = *++p;
    if (yych <= 0x7F)
      goto yy37;
    if (yych <= 0xBF)
      goto yy22;
  yy37:
    p = marker;
    if (yyaccept == 0) {
      goto yy24;
    } else {
      goto yy26;
    }
  yy38:
    yych = *++p;
    if (yych <= 0x9F)
      goto yy37;
    if (yych <= 0xBF)
      goto yy36;
    goto yy37;
  yy39:
    yych = *++p;
    if (yych <= 0x7F)
      goto yy37;
    if (yych <= 0xBF)
      goto yy36;
    goto yy37;
  yy40:
    yych = *++p;
    if (yych <= 0x7F)
      goto yy37;
    if (yych <= 0x9F)
      goto yy36;
    goto yy37;
  yy41:
    yych = *++p;
    if (yych <= 0x8F)
      goto yy37;
    if (yych <= 0xBF)
      goto yy39;
    goto yy37;
  yy42:
    yych = *++p;
    if (yych <= 0x7F)
      goto yy37;
    if (yych <= 0xBF)
      goto yy39;
    goto yy37;
  yy43:
    yych = *++p;
    if (yych <= 0x7F)
      goto yy37;
    if (yych <= 0x8F)
      goto yy39;
    goto yy37;
  }
}

bufsize_t _scan_table_cell_end(const unsigned char *p) {
  const unsigned char *start = p;

  {
    unsigned char yych;
    static const unsigned char yybm[] = {
        0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 128, 128, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   128, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0,
    };
    yych = *p;
    if (yych == '|')
      goto yy48;
    ++p;
    { return 0; }
  yy48:
    yych = *++p;
    if (yybm[0 + yych] & 128) {
      goto yy48;
    }
    { return (bufsize_t)(p - start); }
  }
}

bufsize_t _scan_table_row_end(const unsigned char *p) {
  const unsigned char *marker = NULL;
  const unsigned char *start = p;

  {
    unsigned char yych;
    static const unsigned char yybm[] = {
        0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 128, 128, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   128, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   0, 0,   0,   0, 0, 0,
    };
    yych = *p;
    if (yych <= '\f') {
      if (yych <= 0x08)
        goto yy53;
      if (yych == '\n')
        goto yy56;
      goto yy55;
    } else {
      if (yych <= '\r')
        goto yy58;
      if (yych == ' ')
        goto yy55;
    }
  yy53:
    ++p;
  yy54 : { return 0; }
  yy55:
    yych = *(marker = ++p);
    if (yych <= 0x08)
      goto yy54;
    if (yych <= '\r')
      goto yy60;
    if (yych == ' ')
      goto yy60;
    goto yy54;
  yy56:
    ++p;
    { return (bufsize_t)(p - start); }
  yy58:
    yych = *++p;
    if (yych == '\n')
      goto yy56;
    goto yy54;
  yy59:
    yych = *++p;
  yy60:
    if (yybm[0 + yych] & 128) {
      goto yy59;
    }
    if (yych <= 0x08)
      goto yy61;
    if (yych <= '\n')
      goto yy56;
    if (yych <= '\r')
      goto yy62;
  yy61:
    p = marker;
    goto yy54;
  yy62:
    yych = *++p;
    if (yych == '\n')
      goto yy56;
    goto yy61;
  }
}

bufsize_t _scan_tasklist(const unsigned char *p) {
  const unsigned char *marker = NULL;
  const unsigned char *start = p;

  {
    unsigned char yych;
    static const unsigned char yybm[] = {
        0,   0,   0,   0,   0,   0,   0,   0,   0,   64,  0, 64, 64, 0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        64,  0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 0,  0,  0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
        0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0, 0,  0,  0, 0, 0,
    };
    yych = *p;
    if (yych <= ' ') {
      if (yych <= '\n') {
        if (yych == '\t')
          goto yy67;
      } else {
        if (yych <= '\f')
          goto yy67;
        if (yych >= ' ')
          goto yy67;
      }
    } else {
      if (yych <= ',') {
        if (yych <= ')')
          goto yy65;
        if (yych <= '+')
          goto yy68;
      } else {
        if (yych <= '-')
          goto yy68;
        if (yych <= '/')
          goto yy65;
        if (yych <= '9')
          goto yy69;
      }
    }
  yy65:
    ++p;
  yy66 : { return 0; }
  yy67:
    yych = *(marker = ++p);
    if (yybm[0 + yych] & 64) {
      goto yy70;
    }
    if (yych <= ',') {
      if (yych <= ')')
        goto yy66;
      if (yych <= '+')
        goto yy73;
      goto yy66;
    } else {
      if (yych <= '-')
        goto yy73;
      if (yych <= '/')
        goto yy66;
      if (yych <= '9')
        goto yy74;
      goto yy66;
    }
  yy68:
    yych = *(marker = ++p);
    if (yych <= '\n') {
      if (yych == '\t')
        goto yy75;
      goto yy66;
    } else {
      if (yych <= '\f')
        goto yy75;
      if (yych == ' ')
        goto yy75;
      goto yy66;
    }
  yy69:
    yych = *(marker = ++p);
    if (yych <= 0x1F) {
      if (yych <= '\t') {
        if (yych <= 0x08)
          goto yy78;
        goto yy73;
      } else {
        if (yych <= '\n')
          goto yy66;
        if (yych <= '\f')
          goto yy73;
        goto yy78;
      }
    } else {
      if (yych <= 0x7F) {
        if (yych <= ' ')
          goto yy73;
        goto yy78;
      } else {
        if (yych <= 0xC1)
          goto yy66;
        if (yych <= 0xF4)
          goto yy78;
        goto yy66;
      }
    }
  yy70:
    yych = *++p;
    if (yybm[0 + yych] & 64) {
      goto yy70;
    }
    if (yych <= ',') {
      if (yych <= ')')
        goto yy72;
      if (yych <= '+')
        goto yy73;
    } else {
      if (yych <= '-')
        goto yy73;
      if (yych <= '/')
        goto yy72;
      if (yych <= '9')
        goto yy74;
    }
  yy72:
    p = marker;
    goto yy66;
  yy73:
    yych = *++p;
    if (yych == '[')
      goto yy72;
    goto yy76;
  yy74:
    yych = *++p;
    if (yych <= '\n') {
      if (yych == '\t')
        goto yy73;
      goto yy78;
    } else {
      if (yych <= '\f')
        goto yy73;
      if (yych == ' ')
        goto yy73;
      goto yy78;
    }
  yy75:
    yych = *++p;
  yy76:
    if (yych <= '\f') {
      if (yych == '\t')
        goto yy75;
      if (yych <= '\n')
        goto yy72;
      goto yy75;
    } else {
      if (yych <= ' ') {
        if (yych <= 0x1F)
          goto yy72;
        goto yy75;
      } else {
        if (yych == '[')
          goto yy86;
        goto yy72;
      }
    }
  yy77:
    yych = *++p;
  yy78:
    if (yybm[0 + yych] & 128) {
      goto yy77;
    }
    if (yych <= 0xC1) {
      if (yych <= '\f') {
        if (yych <= 0x08)
          goto yy73;
        if (yych == '\n')
          goto yy72;
        goto yy75;
      } else {
        if (yych == ' ')
          goto yy75;
        if (yych <= 0x7F)
          goto yy73;
        goto yy72;
      }
    } else {
      if (yych <= 0xED) {
        if (yych <= 0xDF)
          goto yy79;
        if (yych <= 0xE0)
          goto yy80;
        if (yych <= 0xEC)
          goto yy81;
        goto yy82;
      } else {
        if (yych <= 0xF0) {
          if (yych <= 0xEF)
            goto yy81;
          goto yy83;
        } else {
          if (yych <= 0xF3)
            goto yy84;
          if (yych <= 0xF4)
            goto yy85;
          goto yy72;
        }
      }
    }
  yy79:
    yych = *++p;
    if (yych <= 0x7F)
      goto yy72;
    if (yych <= 0xBF)
      goto yy73;
    goto yy72;
  yy80:
    yych = *++p;
    if (yych <= 0x9F)
      goto yy72;
    if (yych <= 0xBF)
      goto yy79;
    goto yy72;
  yy81:
    yych = *++p;
    if (yych <= 0x7F)
      goto yy72;
    if (yych <= 0xBF)
      goto yy79;
    goto yy72;
  yy82:
    yych = *++p;
    if (yych <= 0x7F)
      goto yy72;
    if (yych <= 0x9F)
      goto yy79;
    goto yy72;
  yy83:
    yych = *++p;
    if (yych <= 0x8F)
      goto yy72;
    if (yych <= 0xBF)
      goto yy81;
    goto yy72;
  yy84:
    yych = *++p;
    if (yych <= 0x7F)
      goto yy72;
    if (yych <= 0xBF)
      goto yy81;
    goto yy72;
  yy85:
    yych = *++p;
    if (yych <= 0x7F)
      goto yy72;
    if (yych <= 0x8F)
      goto yy81;
    goto yy72;
  yy86:
    yych = *++p;
    if (yych <= 'W') {
      if (yych != ' ')
        goto yy72;
    } else {
      if (yych <= 'X')
        goto yy87;
      if (yych != 'x')
        goto yy72;
    }
  yy87:
    yych = *++p;
    if (yych != ']')
      goto yy72;
    yych = *++p;
    if (yych <= '\n') {
      if (yych != '\t')
        goto yy72;
    } else {
      if (yych <= '\f')
        goto yy89;
      if (yych != ' ')
        goto yy72;
    }
  yy89:
    yych = *++p;
    if (yych <= '\n') {
      if (yych == '\t')
        goto yy89;
    } else {
      if (yych <= '\f')
        goto yy89;
      if (yych == ' ')
        goto yy89;
    }
    { return (bufsize_t)(p - start); }
  }
}
