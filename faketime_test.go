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

func TestSleep(t *testing.T) {
	ft := NewFakeTime()
	if ft == nil {
		t.Fatal("NewFakeTime() returned nil")
	}
	defer ft.Close()
	time.Sleep(time.Second)
	t.Logf("%s\n", time.Now().Format(time.RFC3339))
}
