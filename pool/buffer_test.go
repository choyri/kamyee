package kypool

import (
	"bytes"
	"testing"
)

const TestString = "abcdefghijklmnopqrstuvwxyz"

func TestGetBuffer4096(t *testing.T) {
	pool := GetBuffer4096()

	buf := pool.Get()
	defer pool.Put(buf)

	if c := buf.Cap(); c != 4096 {
		t.Errorf("cap: expect %d, but get %d.", 4096, c)
	}

	buf.WriteString(TestString)
	if s := buf.String(); s != TestString {
		t.Errorf("writeString: expect %s, but get %s.", TestString, s)
	}
}

func BenchmarkGetBuffer4096(b *testing.B) {
	pool := GetBuffer4096()

	for i := 0; i < b.N; i++ {
		var (
			err error
			buf = pool.Get()
		)

		_, err = buf.WriteString(TestString)
		if err != nil {
			b.Errorf("writeString error: %v", err)
		}

		pool.Put(buf)
	}
}

func BenchmarkGetBufferWithSize4096(b *testing.B) {
	pool := GetBuffer(WithSize(4096))

	for i := 0; i < b.N; i++ {
		var (
			err error
			buf = pool.Get()
		)

		_, err = buf.WriteString(TestString)
		if err != nil {
			b.Errorf("writeString error: %v", err)
		}

		pool.Put(buf)
	}
}

func BenchmarkNewBuffer4094(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var (
			err error
			buf = bytes.NewBuffer(make([]byte, 4096))
		)

		_, err = buf.WriteString(TestString)
		if err != nil {
			b.Errorf("writeString error: %v", err)
		}
	}
}
