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
> go.flibuste.net/werr/test.errorf() errorf.go:8
errorf: fail`},
		{wrapf(), `
> go.flibuste.net/werr/test.wrapf() wrapf.go:8
wrapf: fail`},
		{fnew(), `
> go.flibuste.net/werr/test.fnew() new.go:6
new`},
		{ftrace(), `
> go.flibuste.net/werr/test.ftrace() ftrace.go:8
fail`},
		{two(), `
> go.flibuste.net/werr/test.two() two.go:8
two: 
> go.flibuste.net/werr/test.one() two.go:11
one: fail`},
	} {
		res := tst.err.Error()
		if res != tst.res {
			t.Errorf("%d Should be [%s] is [%s]", i, tst.res, res)
		}
	}
}
