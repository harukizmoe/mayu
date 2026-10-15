# Notes for coding agents

Mayu is a Material Design 3 (Material You) component library for MyGo's native
GPU UI, `mygo/ui`. The module is `github.com/harukizmoe/mayu`, the package is
`mayu`, and the repository directory is `mayui`.

## Branches and pull requests

`main` only ever takes merges. Never commit to it, never push to it.

```sh
git checkout -b <type>/<topic>          # feat/, fix/, docs/, chore/
# ... work, one logical change per branch ...
git push -u origin HEAD
gh pr create --title "type(scope): summary" --body-file <file>
gh pr merge --squash --delete-branch
```

Branch protection on `main` rejects a direct push, so a mistake fails rather
than landing. The squash merge gives `main` one commit per change.

Commit messages and PR titles use Conventional Commits, in the imperative:
`fix(segmented): share the track width between segments`.

## Preflight

A pull request must pass all three, and `.github/workflows/ci.yml` runs exactly
them on every pull request and on `main`:

```sh
gofmt -s -l .      # prints nothing
go vet ./...
go test -race ./...
```

## The gallery is the acceptance surface

A component the gallery does not show is not shipped. After adding or changing
one, demonstrate it in `examples/gallery/main.go` and cover it in
`examples/gallery/main_test.go`.

To find a component the gallery never calls:

```sh
grep -oE '^func [A-Z][A-Za-z0-9]*\(' mayu.go | sed 's/func //;s/(//' |
  while read s; do
    grep -q "mayu\.$s\b" examples/gallery/main.go || echo "NOT DEMOED: $s"
  done
```

## Tests

Widgets are tested headlessly with `ui.NewTester`, which builds the view,
delivers real pointer and key events, and paints a frame:

```go
tt := ui.NewTester(func(c *ui.Context) { /* the view */ }, width, height)
tt.Click("Save")          // by the text or label an element shows
tt.Frame()                // after the state the view reads has changed
img := tt.Image()         // the painted frame, for pixel assertions
```

Things worth knowing:

- **Lay out the whole view.** Content below the tester's height gets a zero
  size and silently ignores clicks, so give the gallery tester enough height
  (`examples/gallery/main_test.go` uses `1400x2900`). A click that does nothing
  is usually a click on an element that was never laid out.
- **Assert on behaviour, not on the code.** Prefer a test that fails when the
  behaviour regresses: read a pixel for a colour or a shape, a rect for a
  layout, a clock for a delay. `tt.Find` returns the box of the first element
  showing or labelled with a string, which is the element that drew it.
- **A tooltip has a 600ms delay** before it shows; a test must wait that long.
- `NewTester` returns after settling the first frame. A toast added during that
  build shows on the next one, so call `tt.Frame()` before asserting.

## How the library is put together

| Package | Holds |
|---|---|
| `color` | CAM16 and HCT, tonal palettes, the scheme a seed generates |
| `theme` | the design tokens: `ColorScheme`, `Density`, shape, typography |
| `components` | the widgets, each a builder over a `mygo/ui` base element |
| `symbols` | the Material Symbols the widgets use |
| `mayu.go` | the facade: `Use`, `CurrentTheme`, and one constructor per component |

- **A component takes a `*theme.Theme`**, which its facade wrapper gets from
  `CurrentTheme(c)`. `Use` installs the theme in the window's `ui.Local` at the
  root and also pushes the M3 roles into `mygo/ui`'s own `ui.Theme`, so widgets
  this library does not style still match.
- **Builders embed `ui.Element`**, so a caller can chain the raw element API and
  pass `builder.Element` where an element is wanted — `mayu.Tooltip` and
  `mayu.Menu` take an anchor that way.
- **Build children where the state is known.** A builder that reads state later
  cannot: `Children` runs the closure as it is called, during the build pass.
  Set everything a child draws from before creating it.
- **Style through the state layers.** `ApplyStateLayer` composites Material 3's
  hover, focus and pressed opacities over a container; widgets read
  `Hovered()`, `Focused()` and `Pressed()` as they build.
