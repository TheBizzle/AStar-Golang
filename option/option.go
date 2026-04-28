// Package option defines a data structure for indicating possibly-undefined values, as is common in
// functional programming.
package option

import "reflect"

type Option[T any] interface {
	isOption()
	Get() T
	IsEmpty() bool
}

type None[T any] struct{}

type Some[T any] struct {
	value T
}

func (n None[T]) isOption() {}
func (s Some[T]) isOption() {}

func (n None[T]) Get() T { panic("None.get") }
func (s Some[T]) Get() T { return s.value }

func (n None[T]) IsEmpty() bool { return true }
func (s Some[T]) IsEmpty() bool { return false }

func NewOption[T any](x T) Option[T] {
	v := reflect.ValueOf(any(x))
	if !v.IsValid() || (v.Kind() == reflect.Ptr && v.IsNil()) {
		return None[T]{}
	} else {
		return Some[T]{value: x}
	}
}
