package asserts

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, actual, expected interface{}, msg string) {
	if !reflect.DeepEqual(actual, expected) {
		t.Error(msg)
	}
}

func AssertTrue(t *testing.T, actual bool, msg string) {
	if !actual {
		t.Error(msg)
	}
}

func AssertFalse(t *testing.T, actual bool, msg string) {
	if actual {
		t.Error(msg)
	}
}

func AssertIsNil(t *testing.T, actual interface{}, msg string) {
	if actual != nil {
		t.Error(msg)
	}
}

func AssertIsNotNil(t *testing.T, actual interface{}, msg string) {
	if actual == nil {
		t.Error(msg)
	}
}
