package support

import "time"

const (
	DateDefaultLayout = "2006-01-02"
	TimeDefaultLayout = "2006-01-02 15:04:05"
)

func FormatTime(t time.Time, layout ...string) string {
	l := TimeDefaultLayout

	if len(layout) > 0 && layout[0] != "" {
		l = layout[0]
	}

	return t.Format(l)
}

func ParseTime(value string, layout ...string) (time.Time, error) {
	l := TimeDefaultLayout

	if len(layout) > 0 && layout[0] != "" {
		l = layout[0]
	}

	return time.ParseInLocation(l, value, time.Local)
}
