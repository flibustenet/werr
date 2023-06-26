// Copyright (c) 2023 William Dode. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import (
	"testing"
)

func TestTraces(t *testing.T) {
	for i, tst := range []struct {
		err error
		res string
	}{
		{errorf(), `
> werr/test.errorf() errorf.go:8
errorf: fail`},
		{wrapf(), `
> werr/test.wrapf() wrapf.go:8
wrapf: fail`},
		{fnew(), `
> werr/test.fnew() new.go:6
new`},
		{ftrace(), `
> werr/test.ftrace() ftrace.go:8
fail`},
		{two(), `
> werr/test.two() two.go:8
two: 
> werr/test.one() two.go:11
one: fail`},
	} {
		res := tst.err.Error()
		if res != tst.res {
			t.Errorf("%d Should be [%s] is [%s]", i, tst.res, res)
		}
	}
}
