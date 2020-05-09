package support

import (
	"github.com/choyri/kamyee"
	"os"
	"testing"
)

func TestGetAppEnv(t *testing.T) {
	cases := []struct {
		value string
		want  string
	}{
		{
			want: kamyee.EnvLocal,
		},
		{
			value: kamyee.EnvTesting,
			want:  kamyee.EnvTesting,
		},
	}

	var err error

	for _, v := range cases {
		_ = os.Unsetenv(kamyee.KeyAppEnv)

		if v.value != "" {
			err = os.Setenv(kamyee.KeyAppEnv, v.value)
			if err != nil {
				t.Fatal(err)
			}
		}

		if got := GetAppEnv(); got != v.want {
			t.Errorf("value={%s}, want %s, got {%s}", v.value, v.want, got)
		}
	}
}

func TestGetAppName(t *testing.T) {
	cases := []struct {
		value string
		want  string
	}{
		{
			value: "",
			want:  "",
		},
		{
			value: "hello",
			want:  "hello",
		},
	}

	var err error

	for _, v := range cases {
		_ = os.Unsetenv(kamyee.KeyAppName)

		err = os.Setenv(kamyee.KeyAppName, v.value)
		if err != nil {
			t.Fatal(err)
		}

		if got := GetAppName(); got != v.want {
			t.Errorf("value={%s}, want %s, got {%s}", v.value, v.want, got)
		}
	}
}
