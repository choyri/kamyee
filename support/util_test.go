package support

import "testing"

func TestGetCJK(t *testing.T) {
	cases := []struct {
		value string
		want  string
	}{
		{
			value: "你好 world",
			want:  "你好",
		},
		{
			value: "妳好 world",
			want:  "妳好",
		},
	}

	for _, v := range cases {
		if got := GetCJK(v.value); got != v.want {
			t.Errorf("src=%s, want {%s}, got {%s}", v.value, v.want, got)
		}
	}
}

func TestHTTP2HTTPS(t *testing.T) {
	cases := []struct {
		value string
		want  string
	}{
		{
			value: "http://www.baidu.com",
			want:  "https://www.baidu.com",
		},
		{
			value: "http",
			want:  "http",
		},
	}

	for _, v := range cases {
		if got := HTTP2HTTPS(v.value); got != v.want {
			t.Errorf("src=%s, want {%s}, got {%s}", v.value, v.want, got)
		}
	}
}

func TestSnakeCased(t *testing.T) {
	cases := []struct {
		value string
		want  string
	}{
		{
			value: "ID",
			want:  "i_d",
		},
		{
			value: "HelloWorld",
			want:  "hello_world",
		},
	}

	for _, v := range cases {
		if got := SnakeCased(v.value); got != v.want {
			t.Errorf("src=%s, want {%s}, got {%s}", v.value, v.want, got)
		}
	}
}
