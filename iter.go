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

type iterChan[T any] struct {
	list chan T
	v    T
}

func FromChan[T any](in chan T) Iterator[T] {
	return &iterChan[T]{list: in}
}
func (it *iterChan[T]) Next() bool {
	v, ok := <- it.list
	if ok {
		it.v = v
		return true
	}
	return false
}
func (it *iterChan[T]) Value() T {
	return it.v
}

type iterMap[T comparable, U any] struct {
	m reflect.Value
}

func MapKeys[T comparable, U any](in map[T]U) Iterator[T] {
	return iterMap{m:reflect.ValueOf(in).MapRange()}
}
func (it *iterMap[T,U])Next() bool {
	return it.m.Next()
}
func (it *iterMap[T,U])Value() T {
	return it.m.Key()
}

type iterMapAll[T comparable, U any] struct {
	m reflect.Value
}

func MapAll[T comparable, U any](in map[T]U) Iterator[struct{key T, value U}] {
	return iterMapAll{m:reflect.ValueOf(in).MapRange()}
}
func (it *iterMapAll[T,U])Next() bool {
	return it.m.Next()
}
func (it *iterMapAll[T,U])Value() T {
	return struct{T,U}{it.m.Key(), it.m.Value}
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

// Some does: reduce(), each(), find(), all(), or consume values.
// Ex: Reducer:
// sum := 0
// Some(FromList([]int{1,3,5}, func(i int) {sum+=i; return false})
func Some[T, U any](in Iterator[T], f func(T) (stop bool)) {
	for in.Next() {
		if f(in.Value) {
			return
		}
	}
}

// TODO Py IterTools
// TODO marshal/unmarshal
