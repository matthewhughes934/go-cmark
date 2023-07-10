package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_renderFilesErrorsOnUnreadableFile(t *testing.T) {
	unreadableFile := t.TempDir()
	expectedErrorPrefix := "Failed to read " + unreadableFile

	require.ErrorContains(
		t,
		renderFiles([]string{unreadableFile}, 0),
		expectedErrorPrefix,
	)
}

func Test_renderFilesNoErrorOnReadableFile(t *testing.T) {
	content := "# My file\n"
	markdownFile := filepath.Join(t.TempDir(), "file.md")

	require.NoError(t, os.WriteFile(markdownFile, []byte(content), 0o600))
	// don't bother capturing stdout and asserting on the content
	require.NoError(t, renderFiles([]string{markdownFile}, 80))
}
