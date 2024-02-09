package patherr

import "fmt"

// Wrap wraps the error into a PathError or appends the path to an already
// existing PathError.
func Wrap(err error, subPath ...Element) error {
	if err, ok := err.(*PathError); ok {
		err.Path = append(subPath, err.Path...)
		return err
	}
	return &PathError{Inner: err, Path: subPath}
}

// PathError is an error with a path where the error occured.
type PathError struct {
	Inner error
	Path  Path
}

func (e *PathError) String() string {
	return fmt.Sprintf("error at target path %s: %s", e.Path, e.Inner)
}

func (e *PathError) Unwrap() error { return e.Inner }
func (e *PathError) Error() string { return e.String() }

type Path []Element

func (p Path) String() string {
	if len(p) == 0 {
		return "."
	}

	s := "." + p[0].String()
	for _, element := range p[1:] {
		s += element.Prefix() + element.String()
	}
	return s
}

type Element interface {
	String() string
	Prefix() string
}

type Index int

func (i Index) String() string { return fmt.Sprintf("[%d]", int(i)) }
func (Index) Prefix() string   { return "" }

func Key(key any) Element {
	return KeyElement{key: key}
}

type KeyElement struct{ key any }

func (key KeyElement) String() string {
	return fmt.Sprintf("[%q]", fmt.Sprint(key.key))
}
func (KeyElement) Prefix() string { return "" }

type Field string

func (field Field) String() string { return string(field) }
func (Field) Prefix() string       { return "." }
