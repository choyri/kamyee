package support

import (
	"fmt"
	"os"
	"testing"
)

func TestGetBoolEnv(t *testing.T) {
	cases := []struct {
		value        string
		defaultValue bool
		want         bool
	}{
		{
			value:        "true",
			defaultValue: false,
			want:         true,
		},
		{
			value:        "false",
			defaultValue: true,
			want:         false,
		},
		{
			value:        "1",
			defaultValue: false,
			want:         true,
		},
		{
			value:        "1 ",
			defaultValue: false,
			want:         true,
		},
		{
			value:        "0",
			defaultValue: true,
			want:         false,
		},
		{
			value:        " 0",
			defaultValue: true,
			want:         false,
		},

		{
			value:        "",
			defaultValue: true,
			want:         true,
		},
		{
			value:        "",
			defaultValue: false,
			want:         false,
		},
	}

	var err error

	for k, v := range cases {
		key := fmt.Sprintf("test_bool_env_%d", k)

		if v.value != "" {
			err = os.Setenv(key, v.value)
			if err != nil {
				t.Fatal(err)
			}
		}

		if got := GetBoolEnv(key, v.defaultValue); got != v.want {
			t.Errorf("value=[%s], defaultValue=[%t], want %t, got %t", v.value, v.defaultValue, v.want, got)
		}
	}
}

func TestGetInt64Env(t *testing.T) {
	cases := []struct {
		value        string
		defaultValue int64
		want         int64
	}{
		{
			value:        "-1",
			defaultValue: 0,
			want:         -1,
		},
		{
			value:        "0",
			defaultValue: 1,
			want:         0,
		},
		{
			value:        "1",
			defaultValue: 0,
			want:         1,
		},
		{
			value:        "",
			defaultValue: -1,
			want:         -1,
		},
		{
			value:        "",
			defaultValue: 1,
			want:         1,
		},
		{
			value:        " -1",
			defaultValue: 0,
			want:         -1,
		},
		{
			value:        "1 ",
			defaultValue: 0,
			want:         1,
		},
	}

	var err error

	for k, v := range cases {
		key := fmt.Sprintf("test_int64_env_%d", k)

		if v.value != "" {
			err = os.Setenv(key, v.value)
			if err != nil {
				t.Fatal(err)
			}
		}

		if got := GetInt64Env(key, v.defaultValue); got != v.want {
			t.Errorf("value=[%s], defaultValue=[%d], want %d, got %d", v.value, v.defaultValue, v.want, got)
		}
	}
}

func TestGetStringEnv(t *testing.T) {
	cases := []struct {
		value        string
		defaultValue string
		want         string
	}{
		{
			value:        "yes",
			defaultValue: "no",
			want:         "yes",
		},
		{
			value:        "no",
			defaultValue: "yes",
			want:         "no",
		},
		{
			value:        "",
			defaultValue: "yes",
			want:         "yes",
		},
		{
			value:        "",
			defaultValue: "no",
			want:         "no",
		},
		{
			value:        "yes ",
			defaultValue: "no",
			want:         "yes",
		},
		{
			value:        " no",
			defaultValue: "yes",
			want:         "no",
		},
		{
			value:        " yes no ",
			defaultValue: "yn",
			want:         "yes no",
		},
	}

	var err error

	for k, v := range cases {
		key := fmt.Sprintf("test_string_env_%d", k)

		if v.value != "" {
			err = os.Setenv(key, v.value)
			if err != nil {
				t.Fatal(err)
			}
		}

		if got := GetStringEnv(key, v.defaultValue); got != v.want {
			t.Errorf("value=[%s], defaultValue=[%s], want %s, got %s", v.value, v.defaultValue, v.want, got)
		}
	}
}
