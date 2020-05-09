package support

import "testing"

func TestStr2Int64(t *testing.T) {
	cases := []struct {
		value string
		want  int64
	}{
		{
			value: "-9223372036854775808",
			want:  -9223372036854775808,
		},
		{
			value: "9223372036854775807",
			want:  9223372036854775807,
		},
		{
			value: "hello",
			want:  0,
		},
		{
			value: "",
			want:  0,
		},
	}

	for _, v := range cases {
		if got := Str2Int64(v.value); got != v.want {
			t.Errorf("value=%s, want %d, got %d", v.value, v.want, got)
		}
	}
}

func TestStr2Uint64(t *testing.T) {
	cases := []struct {
		value string
		want  uint64
	}{
		{
			value: "18446744073709551615",
			want:  18446744073709551615,
		},
		{
			value: "hello",
			want:  0,
		},
		{
			value: "",
			want:  0,
		},
	}

	for _, v := range cases {
		if got := Str2Uint64(v.value); got != v.want {
			t.Errorf("value=%s, want %d, got %d", v.value, v.want, got)
		}
	}
}

func TestStr2Int(t *testing.T) {
	cases := []struct {
		value string
		want  int
	}{
		{
			value: "-2147483648",
			want:  -2147483648,
		},
		{
			value: "2147483647",
			want:  2147483647,
		},
		{
			value: "hello",
			want:  0,
		},
		{
			value: "",
			want:  0,
		},
	}

	for _, v := range cases {
		if got := Str2Int(v.value); got != v.want {
			t.Errorf("value=%s, want %d, got %d", v.value, v.want, got)
		}
	}
}

func TestStr2Uint(t *testing.T) {
	cases := []struct {
		value string
		want  uint
	}{
		{
			value: "4294967295",
			want:  4294967295,
		},
		{
			value: "hello",
			want:  0,
		},
		{
			value: "",
			want:  0,
		},
	}

	for _, v := range cases {
		if got := Str2Uint(v.value); got != v.want {
			t.Errorf("value=%s, want %d, got %d", v.value, v.want, got)
		}
	}
}

func TestStr2Bool(t *testing.T) {
	cases := []struct {
		value string
		want  bool
	}{
		{
			value: "true",
			want:  true,
		},
		{
			value: "yes",
			want:  true,
		},
		{
			value: "ok",
			want:  true,
		},
		{
			value: "y",
			want:  true,
		},
		{
			value: "1",
			want:  true,
		},
		{
			value: "true ",
			want:  true,
		},
		{
			value: "",
			want:  false,
		},
		{
			value: "false",
			want:  false,
		},
	}

	for _, v := range cases {
		if got := Str2Bool(v.value); got != v.want {
			t.Errorf("value=%s, want %t, got %t", v.value, v.want, got)
		}
	}
}

func TestInt642Str(t *testing.T) {
	cases := []struct {
		value int64
		want  string
	}{
		{
			value: -9223372036854775808,
			want:  "-9223372036854775808",
		},
		{
			value: 9223372036854775807,
			want:  "9223372036854775807",
		},
		{
			value: 0,
			want:  "0",
		},
	}

	for _, v := range cases {
		if got := Int642Str(v.value); got != v.want {
			t.Errorf("value=%d, want %s, got %s", v.value, v.want, got)
		}
	}
}

func TestUint642Str(t *testing.T) {
	cases := []struct {
		value uint64
		want  string
	}{
		{
			value: 18446744073709551615,
			want:  "18446744073709551615",
		},
		{
			value: 0,
			want:  "0",
		},
	}

	for _, v := range cases {
		if got := Uint642Str(v.value); got != v.want {
			t.Errorf("value=%d, want %s, got %s", v.value, v.want, got)
		}
	}
}

func TestInt2Str(t *testing.T) {
	cases := []struct {
		value int
		want  string
	}{
		{
			value: -2147483648,
			want:  "-2147483648",
		},
		{
			value: 2147483647,
			want:  "2147483647",
		},
		{
			value: 0,
			want:  "0",
		},
	}

	for _, v := range cases {
		if got := Int2Str(v.value); got != v.want {
			t.Errorf("value=%d, want %s, got %s", v.value, v.want, got)
		}
	}
}

func TestUint2Str(t *testing.T) {
	cases := []struct {
		value uint
		want  string
	}{
		{
			value: 4294967295,
			want:  "4294967295",
		},
		{
			value: 0,
			want:  "0",
		},
	}

	for _, v := range cases {
		if got := Uint2Str(v.value); got != v.want {
			t.Errorf("value=%d, want %s, got %s", v.value, v.want, got)
		}
	}
}
