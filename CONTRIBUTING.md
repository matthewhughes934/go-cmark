# Contributing

## Set up

As well as a Go installation you will need:

  - A C compiler
  - [pre-commit](https://pre-commit.com/#install)

Testing is this done with:

    go test ./...

And linting via:

    pre-commit run --all-files

## Memory Management

Rather than require users to manage memory allocated by the C library (e.g.from
`cmark_parser_new`) this is handled via finalizers on these objects. The idea is
to present a more familiar Go interface (i.e. one where you don't have to worry
about memory management) though this may need to be rethought if any issues are
encountered with the finalizers.
