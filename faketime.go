package gofaketime

/*
   #include <time.h>
*/
import "C"

import (
	"time"

	"bou.ke/monkey"
)

/*
通过猴子补丁替换time.Now()的方式来支持faketime;
*/
func fakeTime() time.Time {
	return time.Unix(int64(C.time(nil)), 0)
}

type FakeTime struct {
	faker *monkey.PatchGuard
}

func NewFakeTime() *FakeTime {
	return &FakeTime{faker: monkey.Patch(time.Now, fakeTime)}
}

func (f *FakeTime) Close() {
	f.faker.Unpatch()
}

func (f *FakeTime) Restore() {
	f.faker.Restore()
}
