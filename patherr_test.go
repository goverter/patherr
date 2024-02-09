package patherr

import (
	"errors"
	"fmt"
	"testing"
)

var err = errors.New("err")

func TestString(t *testing.T) {

	data := []struct {
		err      error
		expected string
	}{
		{Wrap(err), "error at target path .: err"},
		{Wrap(err, Field("Field")), "error at target path .Field: err"},
		{Wrap(err, Index(1)), "error at target path .[1]: err"},
		{Wrap(err, Key("key")), `error at target path .["key"]: err`},
		{Wrap(err, Key(`o[i`)), `error at target path .["o[i"]: err`},
		{Wrap(err, Key(`o]i`)), `error at target path .["o]i"]: err`},
		{Wrap(err, Key(`o]i`)), `error at target path .["o]i"]: err`},
		{Wrap(err, Key(err)), `error at target path .["err"]: err`},
		{Wrap(err, Key(struct{ Name string }{Name: "hello"})), `error at target path .["{hello}"]: err`},
		{Wrap(err, Key([]string{"abc", "def"})), `error at target path .["[abc def]"]: err`},
		{Wrap(err, Key(`te"'s[]t`)), `error at target path .["te\"'s[]t"]: err`},
		{Wrap(err, Field("Field"), Index(5), Key("key"), Field("Target")), `error at target path .Field[5]["key"].Target: err`},
	}

	for _, x := range data {
		t.Run(fmt.Sprintf(x.expected), func(t *testing.T) {
			if x.expected != x.err.Error() {
				t.Errorf("%q != %q", x.expected, x.err.Error())
			}
		})
	}
}

func TestMulti(t *testing.T) {
	single := Wrap(err, Key("jmattheis"), Index(5), Key("abc"))

	multi := Wrap(err, Key("abc"))
	multi = Wrap(multi, Index(5))
	multi = Wrap(multi, Key("jmattheis"))

	if single.Error() != multi.Error() {
		t.Errorf("%q != %q", single, multi)
	}
}

func TestUnwrap(t *testing.T) {
	single := Wrap(err, Key("jmattheis"))

	if !errors.Is(single, err) {
		t.Error("is check failed")
	}
}
