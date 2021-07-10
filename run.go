package run

import (
	"context"
	"os"
	"os/signal"
)

// Run runs given f with a context created by context.Background and signal.NotifyContext. After f has finished,
// then Run exits the process by calling os.Exit with returned value of f as its exit code.
//
// By default, signals will be passed to the signal.NotifyContext are unix.SIGHUP, unix.SIGINT, unix.SIGTERM and unix.SIGQUIT.
// You can change which signals will Run wait for by passing an Option created by WithSignals.
func Run(f func(context.Context) int, opts ...Option) {
	opt := newOption(opts...)

	os.Exit(func() int {
		ctx, stop := signal.NotifyContext(context.Background(), opt.signals...)
		defer stop()

		return f(ctx)
	}())
}
