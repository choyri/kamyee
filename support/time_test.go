package support

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	cases := []struct {
		t      time.Time
		layout string
		want   string
	}{
		{
			t:    time.Date(2020, 1, 1, 12, 34, 56, 0, time.Local),
			want: "2020-01-01 12:34:56",
		},
		{
			t:      time.Date(2020, 1, 1, 12, 34, 56, 0, time.Local),
			layout: TimeDefaultLayout,
			want:   "2020-01-01 12:34:56",
		},
		{
			t:      time.Date(2020, 1, 1, 12, 34, 56, 0, time.Local),
			layout: DateDefaultLayout,
			want:   "2020-01-01",
		},
	}

	for _, v := range cases {
		if got := FormatTime(v.t, v.layout); got != v.want {
			t.Errorf("t=[%v], layout=[%s], want %s, got %s", v.t, v.layout, v.want, got)
		}
	}
}

func TestParseTime(t *testing.T) {
	cases := []struct {
		value  string
		layout string
		want   time.Time
	}{
		{
			value: "2020-01-01 12:34:56",
			want:  time.Date(2020, 1, 1, 12, 34, 56, 0, time.Local),
		},
		{
			value:  "2020-01-01 12:34:56",
			layout: TimeDefaultLayout,
			want:   time.Date(2020, 1, 1, 12, 34, 56, 0, time.Local),
		},
		{
			value:  "2020-01-01",
			layout: DateDefaultLayout,
			want:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		},
	}

	for _, v := range cases {
		got, err := ParseTime(v.value, v.layout)
		if err != nil {
			t.Fatal(err)
		}

		if got != v.want {
			t.Errorf("value=[%s], layout=[%s], want %v, got %s", v.value, v.layout, v.want, got)
		}
	}
}
