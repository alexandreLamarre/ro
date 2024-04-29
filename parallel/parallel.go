//go:build goexperiment.rangefunc

package parallel

import (
	"context"
	"iter"
	"sync"
)

func Pool[T any](seq []T, workers int) iter.Seq[T] {
	return func(yield func(T) bool) {
		type item struct {
			val T
			i   int
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch := make(chan item, workers)
		var wg sync.WaitGroup
		wg.Add(workers)

		for i := 0; i < workers; i++ {
			go func() {
				defer wg.Done()
				for v := range ch {
					if !yield(v.val) {
						cancel()
						return
					}
				}
			}()
		}

		for i, v := range seq {
			select {
			case ch <- item{val: v, i: i}:
			case <-ctx.Done():
				return
			}
		}

		close(ch)
		wg.Wait()
	}
}
