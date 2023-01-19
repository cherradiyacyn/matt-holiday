package hello

import (
	"testing"
)

func TestGreet(t *testing.T) {
	want := "Hello, test!"
	got := Greet("test")
	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}
