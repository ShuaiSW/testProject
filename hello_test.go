package testProject

import "testing"

func TestHello(t *testing.T) {
	want := "hello world!"
	if got := Hello(); got != want {
		t.Errorf("got:%v; not equal want:%v", got, want)
	}
}
