package string_util

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	// Why doesn't _, case work? why does it need to be a single char?
	for _, e := range cases {
		got := Reverse(e.in)
		if got != e.want {
			t.Errorf("Reverse(%q) == %q, want %q", e.in, got, e.want)
		}
	}
}
