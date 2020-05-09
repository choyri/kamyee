package support

import "testing"

func TestGetBoolPtr(t *testing.T) {
	for _, v := range []bool{true, false} {
		ptr := GetBoolPtr(v)
		if got := *ptr; got != v {
			t.Errorf("want %t, got %t", v, got)
		}
	}
}

func TestGetStringPtr(t *testing.T) {
	for _, v := range []string{"yes", "no"} {
		ptr := GetStringPtr(v)
		if got := *ptr; got != v {
			t.Errorf("want %s, got %s", v, got)
		}
	}
}
