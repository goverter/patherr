# goverter/patherr

<a href="https://github.com/goverter/patherr/actions/workflows/build.yml">
    <img alt="Build Status" src="https://github.com/goverter/patherr/actions/workflows/build.yml/badge.svg">
</a>
 <a href="https://codecov.io/gh/goverter/patherr">
    <img alt="codecov" src="https://codecov.io/gh/goverter/patherr/branch/main/graph/badge.svg">
</a>
<a href="https://goreportcard.com/report/github.com/goverter/patherr">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/goverter/patherr">
</a>
<a href="https://pkg.go.dev/github.com/goverter/patherr">
    <img alt="Go Reference" src="https://pkg.go.dev/badge/github.com/goverter/patherr.svg">
</a>
<a href="https://github.com/goverter/patherr/releases/latest">
    <img alt="latest release" src="https://img.shields.io/github/release/goverter/patherr.svg">
</a>

`github.com/goverter/patherr` provides an implementation for the goverter
feature [`wrapErrorsUsing`][wrapErrorsUsing]. You can use directly or as a
template for creating your own implementation.

Use it by configuring [`wrapErrorsUsing`][wrapErrorsUsing] on your conversion
interface:

```go
// goverter:converter
// goverter:wrapErrorsUsing github.com/goverter/patherr
type Converter interface {
	Convert(Input) (Output, error)
}
```

Report bugs and request features in the
[github.com/jmattheis/goverter](https://github.com/jmattheis/goverter)
repository.

[wrapErrorsUsing]: https://goverter.jmattheis.de/reference/wrapErrorsUsing
