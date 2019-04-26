package golang_simple_module

import "testing"

func TestHello(t *testing.T) {
	want := "Hallo Welt."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
