package run

import (
	"context"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"
)

var defaultSignals = []os.Signal{unix.SIGTERM, unix.SIGINT}

func Run(f func(context.Context) int, opts ...Option) {
	o := &option{}
	for _, opt := range opts {
		opt.apply(o)
	}

	if len(o.signals) == 0 {
		o.signals = defaultSignals
	}

	os.Exit(func() int {
		ctx, stop := signal.NotifyContext(context.Background(), o.signals...)
		defer stop()

		return f(ctx)
	}())
}
