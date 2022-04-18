package optional

import (
	"fmt"
	"reflect"
)

// Some is a value that represents the presence of a value.
func Some[T any](v T) *Option[T] {
	return &Option[T]{value: v}
}

// None is a value that represents the absence of a value.
func None[T any]() *Option[T] {
	return &Option[T]{null: true}
}

// Option is a value that represents the presence of a value or the absence of a value.
type Option[T any] struct {
	value T
	null  bool
}

// Value returns the value of the Option.
func (o *Option[T]) Value() T {
	return o.value
}

// IsNone returns true if the Option is None.
func (o Option[T]) IsNone() bool {
	if o.null {
		return true
	}
	value := reflect.ValueOf(o.value)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return value.IsNil()
	}
	return false
}

// IsSome returns true if the Option is not None.
func (o Option[T]) IsSome() bool {
	return !o.IsNone()
}

// Except returns the value of the Option.
// If the Option is empty, it will panic with given message.
func (o Option[T]) Except(v any) T {
	if !o.IsNone() {
		return o.value
	}
	panic(v)
}

// Unwrap returns the value of the Option.
// If the Option is empty, it will panic.
func (o Option[T]) Unwrap() T {
	return o.Except("called `Unwrap()` on a nil value")
}

// UnwrapOr returns the value of the Option if it is not None, otherwise returns the given value.
func (o Option[T]) UnwrapOr(v T) T {
	if !o.IsNone() {
		return o.value
	}
	return v
}

// UnwrapOrElse returns the value of the Option if it is not None,
// otherwise returns the result of calling the given function.
func (o Option[T]) UnwrapOrElse(f func() T) T {
	if !o.IsNone() {
		return o.value
	}
	return f()
}

// And returns the Option if it is not None, otherwise returns the given Option.
func (o *Option[T]) And(v *Option[T]) *Option[T] {
	if o.IsNone() || v.IsNone() {
		return o
	}
	return v
}

// AndThen returns the Option if it is not None, otherwise returns the given function result.
func (o *Option[T]) AndThen(f func(T) *Option[T]) *Option[T] {
	if o.IsNone() {
		return o
	}
	return f(o.value)
}

// Or returns the Option if it is not None, otherwise returns the given Option.
func (o *Option[T]) Or(v *Option[T]) *Option[T] {
	if o.IsNone() {
		return v
	}
	return o
}

// OrElse returns the Option if it is not None, otherwise returns the given function result.
func (o *Option[T]) OrElse(next func() *Option[T]) *Option[T] {
	if o.IsNone() {
		return next()
	}
	return o
}

// Filter returns the Option if the given predicate returns true, otherwise returns None.
func (o *Option[T]) Filter(f func(T) bool) *Option[T] {
	if f(o.value) {
		return o
	}
	return None[T]()
}

// Map changes the value of the Option with the given function.
// If the Option is None, it will return None.
func (o *Option[T]) Map(f func(T) T) *Option[T] {
	if o.IsNone() {
		return o
	}
	return Some(f(o.value))
}

// MapOr returns the result of the given function if the Option is not empty,
func (o Option[T]) MapOr(v T, f func(T) T) T {
	if o.IsNone() {
		return v
	}
	return f(o.value)
}

// MapOrElse returns the result of the given function if the Option is not empty,
func (o Option[T]) MapOrElse(def func() T, next func(T) T) T {
	if o.IsNone() {
		return def()
	}
	return next(o.value)
}

func (o *Option[T]) Replace(v T) *Option[T] {
	old := o.value
	o.value = v
	return Some(old)
}

// String returns the string representation of the Option.
// If the Option is empty, it will return "None".
func (o Option[T]) String() string {
	if !o.IsNone() {
		return fmt.Sprintf("%v", o.value)
	}
	return "None"
}
