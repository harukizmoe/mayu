## What

<!-- What this changes, in a sentence or two. -->

## Why

<!-- The problem it solves, or the feedback it answers. -->

## How

<!-- The approach, and anything a reviewer would question. -->

## Checks

- [ ] `gofmt -s -l .` prints nothing
- [ ] `go vet ./...` passes
- [ ] `go test -race ./...` passes
- [ ] Every component this adds or changes is shown in `examples/gallery`, with a test that covers it
- [ ] The gallery runs: `cd examples/gallery && go run main.go`
- [ ] Docs updated: the README catalog, and the package comments of new exported API

## Evidence

<!-- For a UI change, what you saw it do: the measured rect, the sampled pixel, the accessibility tree. -->
