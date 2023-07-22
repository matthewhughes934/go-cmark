package gfm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtensions(t *testing.T) {
	for _, tc := range []struct {
		content       string
		renderOpts    *RenderOpts
		extensionName string
		expected      string
	}{
		{
			"|my table|\n|-|\n|blah|\n",
			NewRenderOpts(),
			"table",
			"<table>\n<thead>\n<tr>\n<th>my table</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>blah</td>\n</tr>\n</tbody>\n</table>\n",
		},
		{
			"hello ~~world~~\n",
			NewRenderOpts(),
			"strikethrough",
			"<p>hello <del>world</del></p>\n",
		},
		{
			"visit https://www.github.com",
			NewRenderOpts(),
			"autolink",
			`<p>visit <a href="https://www.github.com">https://www.github.com</a></p>` + "\n",
		},
		{
			// I need to cmark_render_html with the list of extensions
			// cmark_render_html_with_mem(document, options, parser->syntax_extensions, mem);
			"<script>alert(1)</script>\n",
			NewRenderOpts().WithUnsafe(),
			"tagfilter",
			"&lt;script>alert(1)&lt;/script>\n",
		},
		{
			"- [x] done\n- [ ] not done\n",
			NewRenderOpts(),
			"tasklist",
			"<ul>\n<li><input type=\"checkbox\" checked=\"\" disabled=\"\" /> done</li>\n<li><input type=\"checkbox\" disabled=\"\" /> not done</li>\n</ul>\n",
		},
	} {
		t.Run(tc.extensionName, func(t *testing.T) {
			parser := NewParser()
			extension := FindSyntaxExtension(tc.extensionName)
			require.NotNilf(t, extension, "Failed to find extension: %s", tc.extensionName)
			parser.AttachSyntaxExtension(extension)
			document := parser.ParseDocument(tc.content)

			require.Equal(t, tc.expected, RenderHTML(document, tc.renderOpts, parser.SyntaxExtensions()))
		})
	}
}

func TestFindSyntaxExtensionNoExtension(t *testing.T) {
	require.Nil(t, FindSyntaxExtension("not-a-registered-extension"))
}
