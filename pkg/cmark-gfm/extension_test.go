package gfm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtensions(t *testing.T) {
	for _, tc := range []struct {
		content       string
		renderOpts    *RenderOpts
		extensionName ExtensionName
		expected      string
	}{
		{
			"|my table|\n|-|\n|blah|\n",
			NewRenderOpts(),
			Table,
			"<table>\n<thead>\n<tr>\n<th>my table</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>blah</td>\n</tr>\n</tbody>\n</table>\n",
		},
		{
			"hello ~~world~~\n",
			NewRenderOpts(),
			Strikethrough,
			"<p>hello <del>world</del></p>\n",
		},
		{
			"visit https://www.github.com",
			NewRenderOpts(),
			AutoLink,
			`<p>visit <a href="https://www.github.com">https://www.github.com</a></p>` + "\n",
		},
		{
			// I need to cmark_render_html with the list of extensions
			// cmark_render_html_with_mem(document, options, parser->syntax_extensions, mem);
			"<script>alert(1)</script>\n",
			NewRenderOpts().WithUnsafe(),
			Tagfilter,
			"&lt;script>alert(1)&lt;/script>\n",
		},
		{
			"- [x] done\n- [ ] not done\n",
			NewRenderOpts(),
			Tasklist,
			"<ul>\n<li><input type=\"checkbox\" checked=\"\" disabled=\"\" /> done</li>\n<li><input type=\"checkbox\" disabled=\"\" /> not done</li>\n</ul>\n",
		},
	} {
		t.Run(string(tc.extensionName), func(t *testing.T) {
			extension := FindSyntaxExtension(tc.extensionName)
			parser := NewParser().WithSyntaxExtension(extension)
			document := parser.ParseDocument(tc.content)

			require.Equal(
				t,
				tc.expected,
				RenderHTML(document, tc.renderOpts, parser.SyntaxExtensions()),
			)
		})
	}
}
