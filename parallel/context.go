//go:build goexperiment.rangefunc

package parallel

import (
	"context"
	"iter"

	"github.com/alexandreLamarre/ro"
)

// Context returns an iterator that embeds a context with another iterator
func Context[T any](ctx context.Context, seq iter.Seq[T]) ro.SeqCtx[T] {
	return func(yield func(context.Context, T) bool) {
		seq(func(v T) bool {
			return yield(ctx, v)
		})
	}
}

func Timeout[T any](seq ro.SeqCtx[T]) ro.SeqErr[T] {
	return nil
}

func Cancel[T any](seq ro.SeqCtx[T]) ro.SeqErr[T] {
	return nil
}

func ApplyCtx[U, V any](
	seq iter.Seq[U],
	f func(context.Context, U) (V, error),
) ro.SeqErr[V] {
	return nil
}
