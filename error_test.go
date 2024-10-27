package fig

import (
	"errors"
	"testing"
)

func Test_fieldErrors_Error(t *testing.T) {
	fe := make(fieldErrors)

	fe["B"] = errors.New("berr")
	fe["A"] = errors.New("aerr")

	got := fe.Error()

	if want := "A: aerr, B: berr"; want != got {
		t.Fatalf("want %q, got %q", want, got)
	}

	fe = make(fieldErrors)
	got = fe.Error()

	if got != "" {
		t.Fatalf("empty errors returned non-empty string: %s", got)
	}
}
