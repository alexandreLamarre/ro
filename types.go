package ro

import "iter"

func empty[E any]() iter.Seq[E] {
	return func(yield func(E) bool) {}
}
