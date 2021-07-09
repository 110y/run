package run

import (
	"context"
	"os"
	"os/signal"
)

func Run(f func(context.Context) int, opts ...Option) {
	opt := newOption(opts...)

	os.Exit(func() int {
		ctx, stop := signal.NotifyContext(context.Background(), opt.signals...)
		defer stop()

		return f(ctx)
	}())
}
