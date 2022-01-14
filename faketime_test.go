package gofaketime

import (
	"testing"
	"time"
)

func TestNewFakeTime(t *testing.T) {
	ft := NewFakeTime()
	if ft == nil {
		t.Error("NewFakeTime() returned nil")
	}
	defer ft.Close()
	println(time.Now().Format(time.RFC3339))
}
