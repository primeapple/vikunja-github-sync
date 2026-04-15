package utils

import (
	"errors"
	"reflect"
	"testing"
)

func AssertDeepEquals(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v expected %v", got, want)
	}
}

func AssertNotDeepEquals(t *testing.T, got, notWant any) {
	t.Helper()
	if reflect.DeepEqual(got, notWant) {
		t.Errorf("got %v expected to be different from %v", got, notWant)
	}
}

func AssertEquals(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}
}

func AssertIsPositive(t *testing.T, got int) {
	t.Helper()
	if got <= 0 {
		t.Errorf("got %d but wanted positive value", got)
	}
}

func AssertIsNegative(t *testing.T, got int) {
	t.Helper()
	if got >= 0 {
		t.Errorf("got %d but wanted negative value", got)
	}
}

func AssertNil(t *testing.T, got any) {
	t.Helper()
	if got == nil || reflect.ValueOf(got).IsNil() {
		return
	}
	t.Errorf("got %v but wanted nil", got)
}

func AssertNotNil(t *testing.T, got any) {
	t.Helper()
	if got == nil || reflect.ValueOf(got).IsNil() {
		t.Errorf("got nil but wanted non nil")
	}
}

func AssertError(t *testing.T, got, want error) {
	t.Helper()
	AssertNotNil(t, got)
	if !errors.Is(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("expected passed value to be false, but was true")
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("expected passed value to be true, but was false")
	}
}
