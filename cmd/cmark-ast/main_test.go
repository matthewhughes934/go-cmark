package main

import (
	"io"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func runTest(t *testing.T, args ...string) (string, error) {
	t.Helper()
	origStdout := os.Stdout
	origArgs := os.Args

	defer func() {
		os.Stdout = origStdout
		os.Args = origArgs
	}()

	outReader, outWriter, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = outWriter

	os.Args = args
	runErr := cmarkAST(args)

	outWriter.Close()
	stdOut, err := io.ReadAll(outReader)
	require.NoError(t, err)

	return string(stdOut), runErr
}

func TestHappyPath(t *testing.T) {
	for _, tc := range []struct {
		desc        string
		content     string
		expectedOut string
	}{
		{
			"heading only",
			"# My File",
			`(document)
	(heading)
		(text: 'My File')
`,
		},
		{
			"paragraphs",
			"this is a paragraph\n\nso is _this_\n",
			`(document)
	(paragraph)
		(text: 'this is a paragraph')
	(paragraph)
		(text: 'so is ')
		(emph)
			(text: 'this')
`,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			markdownFile := path.Join(t.TempDir(), "file.md")
			require.NoError(t, os.WriteFile(markdownFile, []byte(tc.content), 0o600))

			gotOut, err := runTest(t, "cmarkast", markdownFile)
			require.Equal(t, tc.expectedOut, gotOut)
			require.NoError(t, err)
		})
	}
}

func TestErrorsOnUnreadableFile(t *testing.T) {
	markdownPath := t.TempDir()
	expectedErrorPrefix := "Error reading file " + markdownPath

	gotOut, err := runTest(t, "cmarktest", markdownPath)

	require.Equal(t, "", gotOut)
	require.ErrorContains(t, err, expectedErrorPrefix)
}
