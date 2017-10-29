package errors_test

import (
	. "errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func level_1() error {
	return Wrap(fmt.Errorf(``))
}

func level_2_0() error {
	return level_2_1()
}

func level_2_1() error {
	return Wrap(fmt.Errorf(``))
}

func Test_level_1(t *testing.T) {
	err := level_1()
	expected := `errors/z_test.go:12`
	t.Log(err)
	if test_path(err.Error()) != expected {
		t.Log("expected:", expected, "vs", err)
		t.Fail()
	}
}

func Test_level_2(t *testing.T) {
	err := level_2_0()
	expected := `errors/z_test.go:20`
	t.Log(err)
	if test_path(err.Error()) != expected {
		t.Log("expected:", expected, "vs", err)
		t.Fail()
	}
}

func Test_anonymous(t *testing.T) {
	f2 := func() error {
		return Wrap(fmt.Errorf(``))
	}
	f1 := func() error {
		return f2()
	}
	err := test_path(f1().Error())
	expected := `errors/z_test.go:45`
	t.Log(err)
	if err != expected {
		t.Log("expected:", expected, "vs", err)
		t.Fail()
	}
}

func Test_nil(t *testing.T) {
	err := Wrap(error(nil))
	t.Log(err)
	if err != nil {
		t.Log("expected:", nil, "vs", err)
	}
}

var cause_error = `cause error`

func test_cause() error {
	return Wrap(fmt.Errorf(cause_error))
}

func Test_cause(t *testing.T) {
	err := test_cause()
	t.Log(Cause(err))
	if Cause(err).Error() != cause_error {
		t.Log("expected:", cause_error, "vs", err)
		t.Fail()
	}
}

func test_path(s string) string {
	sep := string(os.PathSeparator)
	a := strings.Split(s, sep)
	if 2 <= len(a) {
		return strings.Join(a[len(a)-2:], sep)
	}
	return s
}
