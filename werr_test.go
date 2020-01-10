package werr

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

var MyErr = errors.New("MyError")

func aA() error {
	return aB()
}

func aB() error {
	return Wrap(MyErr)
}

func reLines(t *testing.T, lines string, resRe string) bool {
	res := []string{}
	for _, r := range strings.Split(resRe, "\n") {
		r = strings.TrimSpace(r)
		if len(r) == 0 {
			continue
		}
		res = append(res, r)
	}
	ok := true
	i := 0
	for j, line := range strings.Split(lines, "\n") {
		if line == "" {
			continue
		}
		if i >= len(res) {
			t.Errorf("Missing regexp %d %#v %d", i, res, len(res))
			ok = false
			continue
		}
		re := res[i]
		i++
		m, err := regexp.MatchString(re, line)
		if err != nil {
			t.Error(err)
			ok = false
		}
		if !m {
			t.Errorf("Error line %d : %s (%s)", j, line, re)
			ok = false
		}
	}
	if !ok {
		t.Errorf("Errors for\n%s", lines)
	}
	return ok
}

// test simple a call b
// return stack with a and b
func Test_WrapAB(t *testing.T) {
	res := SprintSkip(aA(), "Test_WrapAB")
	reLines(t, res, `
		werr_test.go:\d+ \| go.flib.fr/werr.aA
		werr_test.go:\d+ \| go.flib.fr/werr.aB
		MyError`)
}

// first error tested with Is
func Test_WrapIs(t *testing.T) {
	err := aA()
	if !errors.Is(err, MyErr) {
		t.Error("Is should be MyErr")
	}
}

//
// test stack with fmt.Errorf between
//

func bA() error {
	return bB()
}
func bB() error {
	return fmt.Errorf("by fmt.Errorf : %w", bC())
}
func bC() error {
	return Wrap(MyErr)
}

func Test_ByFmt(t *testing.T) {
	res := SprintSkip(bA(), "Test_ByFmt")
	reLines(t, res, `
		werr_test.go:\d+ \| go.flib.fr/werr.bA
		werr_test.go:\d+ \| go.flib.fr/werr.bB
		werr_test.go:\d+ \| go.flib.fr/werr.bC
		by fmt.Errorf : MyError`)
}

//
// test new with werr.Errorf
//
func cA() error {
	return cB()
}
func cB() error {
	return New("New")
}
func Test_New(t *testing.T) {
	res := SprintSkip(cA(), "Test_New")
	reLines(t, res, `
		werr_test.go:\d+ \| go.flib.fr/werr.cA
		werr_test.go:\d+ \| go.flib.fr/werr.cB
		werr.go:\d+\s+\| go.flib.fr/werr.New
		New`)
}

//
// test new with werr.New
//
func dA() error {
	return dB()
}
func dB() error {
	return Errorf("New %s", "here")
}
func Test_Errorf(t *testing.T) {
	res := SprintSkip(dA(), "Test_Errorf")
	reLines(t, res, `
		werr_test.go:\d+ \| go.flib.fr/werr.dA
		werr_test.go:\d+ \| go.flib.fr/werr.dB
		werr.go:\d+\s+\| go.flib.fr/werr.Errorf
		New here`)
}
