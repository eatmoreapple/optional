package optional

import "testing"

func TestOption_IsNone(t *testing.T) {
	a := Some(1)
	if a.IsNone() {
		t.Error("a.IsNone() should be false")
	}
	b := None[any]()
	if !b.IsNone() {
		t.Error("b.IsNone() should be true")
	}
}

func TestOption_IsSome(t *testing.T) {
	a := Some(1)
	if !a.IsSome() {
		t.Error("a.IsSome() should be true")
	}
	b := None[any]()
	if b.IsSome() {
		t.Error("b.IsSome() should be false")
	}
}

func TestOption_IsEmpty(t *testing.T) {
	a := Some[any](1)
	if a.IsNone() {
		t.Error("a is not empty")
	}
	t.Log("a is not empty")

	t.Log("None is empty")
}

func TestOption_Unwrap(t *testing.T) {
	a := Some[int](1)
	if a.Unwrap() != 1 {
		t.Error("a.Unwrap() != 1")
	}
	t.Log("a.Unwrap() == 1")
}

func TestOption_UnwrapOrElse(t *testing.T) {
	a := Some[int](1)
	if a.UnwrapOrElse(func() int { return 2 }) != 1 {
		t.Error("a.UnwrapOrElse(func() int { return 2 }) != 1")
	}

	c := Some[any](nil)
	if c.UnwrapOrElse(func() any { return 2 }) != 2 {
		t.Error("c.UnwrapOrElse(func() int { return 2 }) != 2")
	}
}

func TestOption_And(t *testing.T) {
	a := Some[int](1)
	b := Some[int](2)
	if a.And(b).Unwrap() != 2 {
		t.Error("a.And(b).Unwrap() != 2")
	}
}

func TestOption_Or(t *testing.T) {
	a := Some[int](1)
	b := Some[int](2)
	if a.Or(b).Unwrap() != 1 {
		t.Error("a.Or(b).Unwrap() != 1")
	}
}

func TestOption_Replace(t *testing.T) {
	a := Some[int](1)
	b := a.Replace(2)
	if a.Unwrap() != 2 {
		t.Error("a.Unwrap() != 2")
	}
	if b.Unwrap() != 1 {
		t.Error("b.Unwrap() != 1")
	}
}

func TestOption_Map(t *testing.T) {
	a := Some[int](1)
	b := a.Map(func(i int) int { return i + 1 })
	if a.Unwrap() != 1 {
		t.Error("a.Unwrap() != 1")
	}
	if b.Unwrap() != 2 {
		t.Error("b.Unwrap() != 2")
	}
}

func TestOption_MapOr(t *testing.T) {
	a := Some[int](1)
	if a.MapOr(3, func(i int) int { return i + 1 }) != 2 {
		t.Error("a.MapOr(3, func(i int) int { return i + 1 }) != 2")
	}
	b := Some[any](nil)
	if b.MapOr(3, func(i any) any { return i }) != 3 {
		t.Error("b.MapOr(3, func(i int) int { return i + 1 }) != 3")
	}
}

func TestOption_MapOrElse(t *testing.T) {
	a := Some[int](1)
	if a.MapOrElse(func() int { return 3 }, func(i int) int { return i + 1 }) != 2 {
		t.Error("a.MapOrElse(func() int { return 3 }, func(i int) int { return i + 1 }) != 2")
	}
	b := Some[any](nil)
	if b.MapOrElse(func() any { return 3 }, func(i any) any { return i }) != 3 {
		t.Error("b.MapOrElse(func() int { return 3 }, func(i int) int { return i + 1 }) != 3")
	}
}
