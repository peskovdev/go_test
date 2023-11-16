package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if err != nil || !want.MatchString(msg) {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if err == nil || msg != "" {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHellos(t *testing.T) {
	names := []string{"Bob", "Alice", "Jenny"}

	messages, err := Hellos(names)
	if err != nil {
		t.Fatalf(`Hellos({"Bob", "Alice", "Jenny"}) = %v, want correct map, error`, err)
	}

	for name, msg := range messages {
		want := regexp.MustCompile(`\b` + name + `\b`)
		if !want.MatchString(msg) {
			t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
		}
	}
}
