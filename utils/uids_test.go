package utils

import (
	"testing"
)

func TestNewUUIDString(t *testing.T) {
	u1, e1 := NewUUIDString()
	u2, e2 := NewUUIDString()
	if e1 != nil || e2 != nil {
		t.Errorf("error not nil: e1: %v, e2: %v", e1, e2)
	}

	if u1 == u2 {
		t.Errorf("u1 equals e2: %v", e1)
	}
}
