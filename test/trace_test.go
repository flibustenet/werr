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
> test.errorf(11): errorf: fail`},
		{wrapf(), `
> test.wrapf(11): wrapf: fail`},
		{fnew(), `
> test.fnew(9): new`},
		{wrap(), `
> test.wrap(11): fail`},
		{two(), `
> test.two(14): two: 
> test.one(17): one: fail`},
		{three(), `
> test.three(11): three: 
> test.two(14): two: 
> test.one(17): one: fail`},
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
