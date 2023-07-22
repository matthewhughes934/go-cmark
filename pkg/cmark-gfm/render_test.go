package gfm

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
			document := NewParser().ParseDocument(tc.content)
			require.Equal(t, tc.expected, RenderCommonMark(document, NewRenderOpts(), tc.width))
		})
	}
}

func TestRenderHTML(t *testing.T) {
	for _, tc := range []struct {
		content  string
		opts     *RenderOpts
		expected string
	}{
		{
			"plain paragraph\nwith two lines\n",
			NewRenderOpts(),
			"<p>plain paragraph\nwith two lines</p>\n",
		},
		{
			"plain paragraph\n",
			NewRenderOpts().WithSourcePos(),
			`<p data-sourcepos="1:1-1:15">plain paragraph</p>` + "\n",
		},
		{
			"plain paragraph\nwith two lines\n",
			NewRenderOpts().WithHarbreaks(),
			"<p>plain paragraph<br />\nwith two lines</p>\n",
		},
		{
			"<img>example.com/img.jpg</img>\n",
			NewRenderOpts(),
			"<p><!-- raw HTML omitted -->example.com/img.jpg<!-- raw HTML omitted --></p>\n",
		},
		{
			"<img>example.com/img.jpg</img>\n",
			NewRenderOpts().WithUnsafe(),
			"<p><img>example.com/img.jpg</img></p>\n",
		},
		{
			"plain paragraph\nwith two lines",
			NewRenderOpts().WithNoBreaks(),
			"<p>plain paragraph with two lines</p>\n",
		},
		{
			"```python\nprint ('hello, world!)\n```\n",
			NewRenderOpts().WithGithubPreLang(),
			`<pre lang="python"><code>print ('hello, world!)` + "\n</code></pre>" + "\n",
		},
		{
			"```python version=3.10\nprint ('hello, world!)\n```\n",
			NewRenderOpts().WithFullInfoString(),
			`<pre><code class="language-python" data-meta="version=3.10">print ('hello, world!)` + "\n</code></pre>" + "\n",
		},
	} {
		t.Run(tc.content, func(t *testing.T) {
			root := NewParser().ParseDocument(tc.content)

			require.Equal(t, tc.expected, RenderHTML(root, tc.opts, nil))
		})
	}
}
