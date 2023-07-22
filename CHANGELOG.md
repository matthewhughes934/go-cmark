# Changelog

## 0.1.0 - 2023-07-01

Initial release

## 0.2.0 - 2023-07-22

### Changed

  - **Breaking** Rename `NodeTypeFootNoteReference` to
    `NodeTypeFootnoteReference` to be consistent with the casing used elsewhere
  - **Breaking** Rewrite option handling for rendering and parsing
  - **Breaking** Move `ParseDocument` to be a method on parsers and not a
    standalone function

### Added

  - Add `RenderCommonMark` for both `cmark` and `cmark-gfm` to wrap
    `cmark_render_commonmark`
  - Add `cmd/cmark-go` to render Markdown files to CommonMark, this is mostly
    meant as a demonstration of the library
  - Add `GetFenceInfo` for both libraries to wrap `cmark_node_get_fence_info`
  - Add `ParentFootnoteDef` to `cmark-gfm` to wrap
    `cmark_node_parent_footnote_def`
  - Add `RenderHTML` for both libraries to wrap `cmark_render_html`
  - Add `WithGithubPreLang` as render option to `cmark-gfm`, wraping
    `CMARK_OPT_GITHUB_PRE_LANG`
