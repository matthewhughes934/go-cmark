package cmark

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRenderCommonMark(t *testing.T) {
	for desc, tc := range map[string]struct {
		content  string
		width    int
		expected string
	}{
		"empty in empty out": {"\n", 1, "\n"},
		"basic document unchanged": {
			"# My Document\n\nWith a paragraph\n",
			80,
			"# My Document\n\nWith a paragraph\n",
		},
		"wraps long lines": {
			strings.Repeat("a", 10) + " " + strings.Repeat("a", 10) + "\n",
			10,
			strings.Repeat("a", 10) + "\n" + strings.Repeat("a", 10) + "\n",
		},
		"basic reformat": {
			"- a dot point\n",
			80,
			"  - a dot point\n",
		},
	} {
		t.Run(desc, func(t *testing.T) {
			document := ParseDocument(tc.content, ParserOptDefault)
			require.Equal(t, tc.expected, RenderCommonMark(document, ParserOptDefault, tc.width))
		})
	}
}
