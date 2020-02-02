package kypool

type Options struct {
	Size uint
}

type Option func(*Options)

func WithSize(size uint) Option {
	return func(opt *Options) {
		opt.Size = size
	}
}
