package support

import (
	"testing"
)

func TestInStringSlice(t *testing.T) {
	cases := []struct {
		needle   string
		haystack []string
		want     bool
	}{
		{
			needle:   "hello",
			haystack: []string{"hello", "world"},
			want:     true,
		},
		{
			needle:   "hi",
			haystack: []string{"hello", "world"},
			want:     false,
		},
	}

	for _, v := range cases {
		if got := InStringSlice(v.needle, v.haystack); got != v.want {
			t.Errorf("needle=[%s], haystack={%v}, want %t, got %t", v.needle, v.haystack, v.want, got)
		}
	}
}

func TestInInt64Slice(t *testing.T) {
	cases := []struct {
		needle   int64
		haystack []int64
		want     bool
	}{
		{
			needle:   -9223372036854775808,
			haystack: []int64{-9223372036854775808, 9223372036854775807},
			want:     true,
		},
		{
			needle:   0,
			haystack: []int64{-9223372036854775808, 9223372036854775807},
			want:     false,
		},
	}

	for _, v := range cases {
		if got := InInt64Slice(v.needle, v.haystack); got != v.want {
			t.Errorf("needle=[%d], haystack={%v}, want %t, got %t", v.needle, v.haystack, v.want, got)
		}
	}
}

func TestInUint64Slice(t *testing.T) {
	cases := []struct {
		needle   uint64
		haystack []uint64
		want     bool
	}{
		{
			needle:   18446744073709551615,
			haystack: []uint64{0, 18446744073709551615},
			want:     true,
		},
		{
			needle:   18446744073709551614,
			haystack: []uint64{0, 18446744073709551615},
			want:     false,
		},
	}

	for _, v := range cases {
		if got := InUint64Slice(v.needle, v.haystack); got != v.want {
			t.Errorf("needle=[%d], haystack={%v}, want %t, got %t", v.needle, v.haystack, v.want, got)
		}
	}
}

func TestInIntSlice(t *testing.T) {
	cases := []struct {
		needle   int
		haystack []int
		want     bool
	}{
		{
			needle:   2147483647,
			haystack: []int{-2147483648, 2147483647},
			want:     true,
		},
		{
			needle:   -2147483647,
			haystack: []int{-2147483648, 2147483647},
			want:     false,
		},
	}

	for _, v := range cases {
		if got := InIntSlice(v.needle, v.haystack); got != v.want {
			t.Errorf("needle=[%d], haystack={%v}, want %t, got %t", v.needle, v.haystack, v.want, got)
		}
	}
}

func TestUinIntSlice(t *testing.T) {
	cases := []struct {
		needle   uint
		haystack []uint
		want     bool
	}{
		{
			needle:   4294967295,
			haystack: []uint{0, 4294967295},
			want:     true,
		},
		{
			needle:   4294967294,
			haystack: []uint{0, 4294967295},
			want:     false,
		},
	}

	for _, v := range cases {
		if got := InUintSlice(v.needle, v.haystack); got != v.want {
			t.Errorf("needle=[%d], haystack={%v}, want %t, got %t", v.needle, v.haystack, v.want, got)
		}
	}
}

func TestIntersectStringSlice(t *testing.T) {
	cases := []struct {
		a    []string
		b    []string
		want []string
	}{
		{
			a:    []string{"hello", "world"},
			b:    []string{"hello", "world"},
			want: []string{"hello", "world"},
		},
		{
			a:    []string{"hello", "world"},
			b:    []string{"hello"},
			want: []string{"hello"},
		},
		{
			a:    []string{"hello", "world"},
			b:    []string{"hi"},
			want: []string{},
		},
		{
			a:    []string{"hello", "world"},
			b:    []string{},
			want: []string{},
		},
		{
			a:    []string{"hello"},
			b:    []string{"hello", "world"},
			want: []string{"hello"},
		},
	}

	for _, v := range cases {
		got := IntersectStringSlice(v.a, v.b)
		if len(got) != len(v.want) {
			t.Errorf("a={%v}, b={%v}, want {%v}, got {%v}", v.a, v.b, v.want, got)
		}

		for _, vv := range got {
			if InStringSlice(vv, v.want) == false {
				t.Errorf("a={%v}, b={%v}, want {%v}, got {%v}", v.a, v.b, v.want, got)
			}
		}
	}
}

func TestIntersectInt64Slice(t *testing.T) {
	cases := []struct {
		a    []int64
		b    []int64
		want []int64
	}{
		{
			a:    []int64{-9223372036854775808, 9223372036854775807},
			b:    []int64{-9223372036854775808, 9223372036854775807},
			want: []int64{-9223372036854775808, 9223372036854775807},
		},
		{
			a:    []int64{-9223372036854775808, 9223372036854775807},
			b:    []int64{-9223372036854775808},
			want: []int64{-9223372036854775808},
		},
		{
			a:    []int64{-9223372036854775808, 9223372036854775807},
			b:    []int64{-9223372036854775807},
			want: []int64{},
		},
		{
			a:    []int64{-9223372036854775808, 9223372036854775807},
			b:    []int64{},
			want: []int64{},
		},
		{
			a:    []int64{-9223372036854775808},
			b:    []int64{-9223372036854775808, 9223372036854775807},
			want: []int64{-9223372036854775808},
		},
	}

	for _, v := range cases {
		got := IntersectInt64Slice(v.a, v.b)
		if len(got) != len(v.want) {
			t.Errorf("a={%v}, b={%v}, want {%v}, got {%v}", v.a, v.b, v.want, got)
		}

		for _, vv := range got {
			if InInt64Slice(vv, v.want) == false {
				t.Errorf("a={%v}, b={%v}, want {%v}, got {%v}", v.a, v.b, v.want, got)
			}
		}
	}
}

func TestIntersectUint64Slice(t *testing.T) {
	cases := []struct {
		a    []uint64
		b    []uint64
		want []uint64
	}{
		{
			a:    []uint64{0, 18446744073709551615},
			b:    []uint64{0, 18446744073709551615},
			want: []uint64{0, 18446744073709551615},
		},
		{
			a:    []uint64{0, 18446744073709551615},
			b:    []uint64{18446744073709551615},
			want: []uint64{18446744073709551615},
		},
		{
			a:    []uint64{0, 18446744073709551615},
			b:    []uint64{18446744073709551614},
			want: []uint64{},
		},
		{
			a:    []uint64{0, 18446744073709551615},
			b:    []uint64{},
			want: []uint64{},
		},
		{
			a:    []uint64{18446744073709551615},
			b:    []uint64{0, 18446744073709551615},
			want: []uint64{18446744073709551615},
		},
	}

	for _, v := range cases {
		got := IntersectUint64Slice(v.a, v.b)
		if len(got) != len(v.want) {
			t.Errorf("a={%v}, b={%v}, want {%v}, got {%v}", v.a, v.b, v.want, got)
		}

		for _, vv := range got {
			if InUint64Slice(vv, v.want) == false {
				t.Errorf("a={%v}, b={%v}, want {%v}, got {%v}", v.a, v.b, v.want, got)
			}
		}
	}
}

func TestIntersectIntSlice(t *testing.T) {
	cases := []struct {
		a    []int
		b    []int
		want []int
	}{
		{
			a:    []int{-2147483648, 2147483647},
			b:    []int{-2147483648, 2147483647},
			want: []int{-2147483648, 2147483647},
		},
		{
			a:    []int{-2147483648, 2147483647},
			b:    []int{-2147483648},
			want: []int{-2147483648},
		},
		{
			a:    []int{-2147483648, 2147483647},
			b:    []int{-2147483647},
			want: []int{},
		},
		{
			a:    []int{-2147483648, 2147483647},
			b:    []int{},
			want: []int{},
		},
		{
			a:    []int{-2147483648},
			b:    []int{-2147483648, 2147483647},
			want: []int{-2147483648},
		},
	}

	for _, v := range cases {
		got := IntersectIntSlice(v.a, v.b)
		if len(got) != len(v.want) {
			t.Errorf("a={%v}, b={%v}, want {%v}, got {%v}", v.a, v.b, v.want, got)
		}

		for _, vv := range got {
			if InIntSlice(vv, v.want) == false {
				t.Errorf("a={%v}, b={%v}, want {%v}, got {%v}", v.a, v.b, v.want, got)
			}
		}
	}
}

func TestIntersectUintSlice(t *testing.T) {
	cases := []struct {
		a    []uint
		b    []uint
		want []uint
	}{
		{
			a:    []uint{0, 4294967295},
			b:    []uint{0, 4294967295},
			want: []uint{0, 4294967295},
		},
		{
			a:    []uint{0, 4294967295},
			b:    []uint{4294967295},
			want: []uint{4294967295},
		},
		{
			a:    []uint{0, 4294967295},
			b:    []uint{4294967294},
			want: []uint{},
		},
		{
			a:    []uint{0, 4294967295},
			b:    []uint{},
			want: []uint{},
		},
		{
			a:    []uint{4294967295},
			b:    []uint{0, 4294967295},
			want: []uint{4294967295},
		},
	}

	for _, v := range cases {
		got := IntersectUintSlice(v.a, v.b)
		if len(got) != len(v.want) {
			t.Errorf("a={%v}, b={%v}, want {%v}, got {%v}", v.a, v.b, v.want, got)
		}

		for _, vv := range got {
			if InUintSlice(vv, v.want) == false {
				t.Errorf("a={%v}, b={%v}, want {%v}, got {%v}", v.a, v.b, v.want, got)
			}
		}
	}
}

func TestGetAddedAndDeletedInUintSlice(t *testing.T) {
	cases := []struct {
		old     []uint
		new     []uint
		added   []uint
		deleted []uint
	}{
		{
			old:     []uint{0, 4294967295},
			new:     []uint{0, 4294967295},
			added:   []uint{},
			deleted: []uint{},
		},
		{
			old:     []uint{0, 4294967295},
			new:     []uint{0, 4294967294},
			added:   []uint{4294967294},
			deleted: []uint{4294967295},
		},
		{
			old:     []uint{0, 4294967295},
			new:     []uint{1, 4294967294},
			added:   []uint{1, 4294967294},
			deleted: []uint{0, 4294967295},
		},
		{
			old:     []uint{0, 4294967295},
			new:     []uint{},
			added:   []uint{},
			deleted: []uint{0, 4294967295},
		},
		{
			old:     []uint{},
			new:     []uint{0, 4294967295},
			added:   []uint{0, 4294967295},
			deleted: []uint{},
		},
	}

	for _, v := range cases {
		added, deleted := GetAddedAndDeletedInUintSlice(v.old, v.new)

		if len(added) != len(v.added) {
			t.Errorf("old={%v}, new={%v}, want added={%v}, got added={%v}", v.old, v.new, v.added, added)
		}

		for _, vv := range added {
			if InUintSlice(vv, v.added) == false {
				t.Errorf("old={%v}, new={%v}, want added={%v}, got added={%v}", v.old, v.new, v.added, added)
			}
		}

		if len(deleted) != len(v.deleted) {
			t.Errorf("old={%v}, new={%v}, want deleted={%v}, got deleted={%v}", v.old, v.new, v.deleted, deleted)
		}

		for _, vv := range deleted {
			if InUintSlice(vv, v.deleted) == false {
				t.Errorf("old={%v}, new={%v}, want deleted={%v}, got deleted={%v}", v.old, v.new, v.deleted, deleted)
			}
		}
	}
}
