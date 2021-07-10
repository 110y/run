package run

import (
	"os"

	"golang.org/x/sys/unix"
)

var _ Option = (*funcOption)(nil)

var defaultSignals = []os.Signal{unix.SIGHUP, unix.SIGINT, unix.SIGTERM, unix.SIGQUIT}

// Option controls how Run behaves.
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

func newOption(opts ...Option) *option {
	o := &option{}
	for _, opt := range opts {
		opt.apply(o)
	}

	if len(o.signals) == 0 {
		o.signals = defaultSignals
	}

	return o
}

// WithSignals returns an Option which specifies which signals will Run wait for with signal.NotifyContext.
func WithSignals(signals ...os.Signal) Option {
	return newFuncOption(func(o *option) {
		o.signals = make([]os.Signal, len(signals))
		for i, s := range signals {
			o.signals[i] = s
		}
	})
}
