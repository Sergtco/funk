package funk

import (
	"errors"
	"iter"
)

func Map[F, T any](items iter.Seq[F], fn func(F) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range items {
			if !yield(fn(item)) {
				return
			}
		}
	}
}

func Fold[Acc, T any](items iter.Seq[T], acc Acc, fn func(Acc, T) Acc) Acc {
	for item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func fold[Acc, T any](next func() (T, bool), acc Acc, fn func(Acc, T) Acc) Acc {
	for val, ok := next(); ok; val, ok = next() {
		acc = fn(acc, val)
	}
	return acc
}

func Reduce[T any](items iter.Seq[T], fn func(T, T) T) (T, error) {
	next, stop := iter.Pull(items)
	defer stop()
	acc, ok := next()
	if !ok {
		return *new(T), errors.New("Unable to Reduce empty sequence.")
	}
	return fold(next, acc, fn), nil
}
