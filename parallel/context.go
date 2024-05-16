//go:build goexperiment.rangefunc

package parallel

import (
	"context"
	"iter"

	"github.com/alexandreLamarre/ro"
)

type ParallelOptions struct {
	Ctx context.Context
}

func (o *ParallelOptions) Apply(opts ...ParallelOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func defaultParallelOptions() *ParallelOptions {
	return &ParallelOptions{
		Ctx: context.Background(),
	}
}

type ParallelOption func(*ParallelOptions)

func WithContext(ctx context.Context) ParallelOption {
	return func(o *ParallelOptions) {
		o.Ctx = ctx
	}
}

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

func Go[T any](seq iter.Seq[T], workers int, opts ...ParallelOption) iter.Seq[T] {
	options := defaultParallelOptions()
	options.Apply(opts...)
	return func(yield func(T) bool) {

	}
}
