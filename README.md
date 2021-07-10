# run

`run` is a tiny module which makes spinning up and tearing down for a server program written in Go be more easier. See an example below for how to use.

```go
import (
	"context"

	"github.com/110y/run"
)

func main() {
	run.Run(func(ctx context.Context) int {
		// Spin up your server here.

		// This blocks until one of termination signals (unix.SIGHUP, unix.SIGINT, unix.SIGTERM or unix.SIGQUIT) will be passed.
		<-ctx.Done()

		// Tear down your server here.

		// After this function has finished, run.Run exits the process with returned value of this function as its exit code.
		return 0
	})
}
```
