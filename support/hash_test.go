package support

import "testing"

func TestMD5(t *testing.T) {
	cases := []struct {
		data []byte
		want string
	}{
		{
			data: []byte("hello"),
			want: "5d41402abc4b2a76b9719d911017c592",
		},
		{
			data: []byte("world"),
			want: "7d793037a0760186574b0282f2f435e7",
		},
	}

	for _, v := range cases {
		if got := MD5(v.data); got != v.want {
			t.Errorf("data={%s}, want %s, got %s", string(v.data), v.want, got)
		}
	}
}
