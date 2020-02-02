package kypool

import (
	"bytes"
	"sync"
)

type bufferPool sync.Pool

var (
	buffer4096 *sync.Pool
	buffer8192 *sync.Pool

	buffer4096Once sync.Once
	buffer8192Once sync.Once
)

func newBufferPool(size uint) *sync.Pool {
	return &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, size))
		},
	}
}

func GetBuffer(opts ...Option) BufferPool {
	options := Options{}

	if len(opts) == 0 {
		goto directDefault
	}

	for _, opt := range opts {
		opt(&options)
	}

	if options.Size > 0 {
		return (*bufferPool)(newBufferPool(options.Size))
	}

directDefault:
	return GetBuffer4096()
}

func GetBuffer4096() BufferPool {
	buffer4096Once.Do(func() {
		buffer4096 = newBufferPool(4096)
	})

	return (*bufferPool)(buffer4096)
}

func GetBuffer8192() BufferPool {
	buffer8192Once.Do(func() {
		buffer8192 = newBufferPool(8192)
	})

	return (*bufferPool)(buffer8192)
}

type BufferPool interface {
	Get() *bytes.Buffer
	Put(*bytes.Buffer)
}

func (bp *bufferPool) Get() *bytes.Buffer {
	ret := (*sync.Pool)(bp).Get().(*bytes.Buffer)
	ret.Reset()
	return ret
}

func (bp *bufferPool) Put(buf *bytes.Buffer) {
	(*sync.Pool)(bp).Put(buf)
}
