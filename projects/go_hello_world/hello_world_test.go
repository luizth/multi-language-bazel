package go_hello_world

import "testing"

func TestGreeter(t *testing.T) {
	expected := "Hello World!"
	actual := HelloWorld()
	if actual != expected {
		t.Errorf("expected %q bu got %q", expected, actual)
	}
}
