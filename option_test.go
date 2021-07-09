package run

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/sys/unix"
)

func TestWithSignals(t *testing.T) {
	t.Parallel()

	actual := &option{}
	WithSignals(unix.SIGTERM).apply(actual)

	expected := &option{signals: []os.Signal{unix.SIGTERM}}

	if diff := cmp.Diff(actual, expected, cmp.AllowUnexported(option{})); diff != "" {
		t.Errorf("\n(-actual, +expected)\n%s", diff)
	}
}

func TestNewOption(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		opts     []Option
		expected *option
	}{
		"empty options": {
			opts: nil,
			expected: &option{
				signals: defaultSignals,
			},
		},
		"WithSignals": {
			opts: []Option{
				WithSignals(unix.SIGTERM),
			},
			expected: &option{
				signals: []os.Signal{unix.SIGTERM},
			},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := newOption(test.opts...)
			if diff := cmp.Diff(actual, test.expected, cmp.AllowUnexported(option{})); diff != "" {
				t.Errorf("\n(-actual, +expected)\n%s", diff)
			}
		})
	}
}
