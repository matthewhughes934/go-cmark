/*
cmark-go renders files as common mark via [cmark.ParseDocument]

if no files are provided STDIN is read

Usage:

	cmark-go [ --width width ] [ FILE ]...
*/
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/matthewhughes934/go-cmark/pkg/cmark"
	"github.com/urfave/cli/v3"
)

func main() { //go-cov:skip
	var opts struct {
		width int64
	}
	cmd := &cli.Command{
		Name: os.Args[0],
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "width",
				Usage:       "Specify wrap width (default 0 = nowrap)",
				Value:       0,
				Destination: &opts.width,
			},
		},
		Action: func(c *cli.Context) error {
			return renderFiles(c.Args().Slice(), int(opts.width))
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func renderFiles(filenames []string, width int) error {
	if len(filenames) == 0 { //go-cov:skip
		filenames = []string{"/dev/stdin"}
	}
	for _, filename := range filenames {
		content, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("Failed to read %s: %v", filename, err)
		}
		document := cmark.NewParser().ParseDocument(string(content))
		fmt.Print(cmark.RenderCommonMark(document, cmark.NewRenderOpts(), width))
	}
	return nil
}
