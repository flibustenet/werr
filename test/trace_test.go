// Copyright (c) 2023 William Dode. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import (
	"errors"
	"testing"

	"go.flibuste.net/werr"
)

func TestTraces(t *testing.T) {
	for i, tst := range []struct {
		err error
		res string
	}{
		{errorf(), `
> werr/test.errorf() errorf.go:11
errorf: fail`},
		{wrapf(), `
> werr/test.wrapf() wrapf.go:11
wrapf: fail`},
		{fnew(), `
> werr/test.fnew() new.go:9
new`},
		{wrap(), `
> werr/test.wrap() wrap.go:11
fail`},
		{two(), `
> werr/test.two() two.go:14
two: 
> werr/test.one() two.go:17
one: fail`},
		{three(), `
> werr/test.three() two.go:11
three: 
> werr/test.two() two.go:14
two: 
> werr/test.one() two.go:17
one: fail`},
	} {
		res := tst.err.Error()
		if res != tst.res {
			t.Errorf("%d Should be [%s] is [%s]", i, tst.res, res)
		}
	}
}

func TestWraping(t *testing.T) {
	err := errors.New("oups")
	newErr := werr.Wrap(err)
	if !errors.Is(newErr, err) {
		t.Errorf("Wrap should wrap")
	}
	newErr = werr.Wrapf(err, "")
	if !errors.Is(newErr, err) {
		t.Errorf("Wrapf should wrap")
	}
}

func TestNull(t *testing.T) {
	newErr := werr.Wrap(nil)
	if newErr != nil {
		t.Errorf("nil should be nil")
	}
	newErr = werr.Wrapf(nil, "")
	if newErr != nil {
		t.Errorf("nil should be nil")
	}
}
