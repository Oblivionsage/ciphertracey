package version

import "testing"

func TestString(t *testing.T) {
	v := String()
	expected := "v0.1.0"
	if v != expected {
		t.Errorf("version.String() = %q, want %q", v, expected)
	}
}
