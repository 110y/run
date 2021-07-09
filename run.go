package run

import (
	"context"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"
)

func Run(f func(context.Context) int) {
	os.Exit(func() int {
		ctx, stop := signal.NotifyContext(context.Background(), unix.SIGTERM, unix.SIGINT)
		defer stop()

		return f(ctx)
	}())
}
