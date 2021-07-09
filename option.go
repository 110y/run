package run

import "os"

var _ Option = (*funcOption)(nil)

type Option interface {
	apply(*option)
}

type option struct {
	signals []os.Signal
}

type funcOption struct {
	f func(*option)
}

func (fo *funcOption) apply(o *option) {
	fo.f(o)
}

func newFuncOption(f func(*option)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func WithSignals(signals ...os.Signal) Option {
	return newFuncOption(func(o *option) {
		o.signals = make([]os.Signal, len(signals))
		for i, s := range signals {
			o.signals[i] = s
		}
	})
}
