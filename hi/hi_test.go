package hi

import (
	"regexp"
	"testing"
)

func TestEmptyHelloWorld(t *testing.T) {
	name := ""
	msg, err := HelloWorld(name)
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHelloWorld(t *testing.T) {
	name := "gholi"
	want := regexp.MustCompile(`\b` + name + `b`)
	msg, err := HelloWorld(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("gholi") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
