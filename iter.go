package iter

import "fmt"

type Iterator[T any] interface {
	Next() bool
	Value() T
}

type iterList[T any] struct {
	list []T
	v    T
}

func FromList[T any](in []T) Iterator[T] {
	return &iterList[T]{list: in}
}
func (it *iterList[T]) Next() bool {
	if len(it.list) > 0 {
		it.v = it.list[0]
		it.list = it.list[1:]
		return true
	}
	return false
}
func (it *iterList[T]) Value() T {
	return it.v
}

type changer[T, U any] struct {
	f  func(T) U
	in Iterator[T]
	v  U
	ok bool
}

func (it *changer[T, U]) Next() bool {
	ok := it.in.Next()
	if ok {
		it.v = it.f(it.in.Value())
	}
	return ok
}
func (it *changer[T, U]) Value() U {
	return it.v
}

func Change[T, U any](in Iterator[T], f func(T) U) Iterator[U] {
	return &changer[T, U]{f: f, in: in}
}
