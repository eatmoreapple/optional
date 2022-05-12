package optional

import (
	"fmt"
	"reflect"
)

// Some is a value that represents the presence of a value.
func Some[T any](v T) Option[T] {
	return &option[T]{value: v}
}

// None is a value that represents the absence of a value.
func None[T any]() Option[T] {
	return &option[T]{null: true}
}

// Option is a value that represents the presence of a value or the absence of a value.
type Option[T any] interface {
	Value() T
	IsNone() bool
	IsSome() bool
	Except(v any) T
	Unwrap() T
	UnwrapOr(v T) T
	UnwrapOrElse(f func() T) T
	And(v Option[T]) Option[T]
	AndThen(f func(T) Option[T]) Option[T]
	Or(v Option[T]) Option[T]
	OrElse(next func() Option[T]) Option[T]
	Filter(f func(T) bool) Option[T]
	Map(f func(T) T) Option[T]
	MapOr(v T, f func(T) T) T
	MapOrElse(def func() T, next func(T) T) T
	Replace(v T) Option[T]
}

// option is a value that represents the presence of a value or the absence of a value.
type option[T any] struct {
	value T
	null  bool
}

// Value returns the value of the option.
func (o *option[T]) Value() T {
	return o.value
}

// IsNone returns true if the option is None.
func (o option[T]) IsNone() bool {
	if o.null {
		return true
	}
	value := reflect.ValueOf(o.value)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return value.IsNil()
	case reflect.Invalid:
		return true
	}
	return false
}

// IsSome returns true if the option is not None.
func (o option[T]) IsSome() bool {
	return !o.IsNone()
}

// Except returns the value of the option.
// If the option is empty, it will panic with given message.
func (o option[T]) Except(v any) T {
	if !o.IsNone() {
		return o.value
	}
	panic(v)
}

// Unwrap returns the value of the option.
// If the option is empty, it will panic.
func (o option[T]) Unwrap() T {
	return o.Except("called `Unwrap()` on a nil value")
}

// UnwrapOr returns the value of the option if it is not None, otherwise returns the given value.
func (o option[T]) UnwrapOr(v T) T {
	if !o.IsNone() {
		return o.value
	}
	return v
}

// UnwrapOrElse returns the value of the option if it is not None,
// otherwise returns the result of calling the given function.
func (o option[T]) UnwrapOrElse(f func() T) T {
	if !o.IsNone() {
		return o.value
	}
	return f()
}

// And returns the option if it is not None, otherwise returns the given option.
func (o *option[T]) And(v Option[T]) Option[T] {
	if o.IsNone() || v.IsNone() {
		return o
	}
	return v
}

// AndThen returns the option if it is not None, otherwise returns the given function result.
func (o *option[T]) AndThen(f func(T) Option[T]) Option[T] {
	if o.IsNone() {
		return o
	}
	return f(o.value)
}

// Or returns the option if it is not None, otherwise returns the given option.
func (o *option[T]) Or(v Option[T]) Option[T] {
	if o.IsNone() {
		return v
	}
	return o
}

// OrElse returns the option if it is not None, otherwise returns the given function result.
func (o *option[T]) OrElse(next func() Option[T]) Option[T] {
	if o.IsNone() {
		return next()
	}
	return o
}

// Filter returns the option if the given predicate returns true, otherwise returns None.
func (o *option[T]) Filter(f func(T) bool) Option[T] {
	if f(o.value) {
		return o
	}
	return None[T]()
}

// Map changes the value of the option with the given function.
// If the option is None, it will return None.
func (o *option[T]) Map(f func(T) T) Option[T] {
	if o.IsNone() {
		return o
	}
	return Some(f(o.value))
}

// MapOr returns the result of the given function if the option is not empty,
func (o option[T]) MapOr(v T, f func(T) T) T {
	if o.IsNone() {
		return v
	}
	return f(o.value)
}

// MapOrElse returns the result of the given function if the option is not empty,
func (o option[T]) MapOrElse(def func() T, next func(T) T) T {
	if o.IsNone() {
		return def()
	}
	return next(o.value)
}

func (o *option[T]) Replace(v T) Option[T] {
	old := o.value
	o.value = v
	return Some(old)
}

// String returns the string representation of the option.
// If the option is empty, it will return "None".
func (o option[T]) String() string {
	if !o.IsNone() {
		return fmt.Sprintf("%v", o.value)
	}
	return "None"
}
