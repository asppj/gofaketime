package gofaketime

/*
   #include <time.h>
*/
import "C"

import (
	"sync"
	"time"

	"bou.ke/monkey"
)

/*
通过猴子补丁替换time.Now()的方式来支持faketime;
*/

var (
	lockerNow     = sync.Mutex{}
	globalLock    = sync.Mutex{}
	isPatched     = false
	globalFaker   *FakeTime
)

func fakeTime() time.Time {
	lockerNow.Lock()
	defer lockerNow.Unlock()
	return time.Unix(int64(C.time(nil)), 0)
}

type FakeTime struct {
	faker     *monkey.PatchGuard
	isActive  bool
	closeLock sync.Mutex
}

func NewFakeTime() *FakeTime {
	globalLock.Lock()
	defer globalLock.Unlock()

	// 防止重复 patching
	if isPatched {
		// 返回已存在的实例
		if globalFaker != nil {
			return globalFaker
		}
	}

	ft := &FakeTime{
		faker:    monkey.Patch(time.Now, fakeTime),
		isActive: true,
	}

	isPatched = true
	globalFaker = ft

	return ft
}

func (f *FakeTime) Close() {
	f.closeLock.Lock()
	defer f.closeLock.Unlock()

	globalLock.Lock()
	defer globalLock.Unlock()

	if !f.isActive {
		return
	}

	if f.faker != nil {
		f.faker.Unpatch()
	}

	f.isActive = false
	isPatched = false
	globalFaker = nil
}

func (f *FakeTime) Restore() {
	f.closeLock.Lock()
	defer f.closeLock.Unlock()

	if !f.isActive || f.faker == nil {
		return
	}

	f.faker.Restore()
}

/*
第二种使用方式。结合 go build -ldflags "--tags dev"
*/
var faker *FakeTime

func Init() {
	globalLock.Lock()
	defer globalLock.Unlock()

	if faker == nil {
		faker = NewFakeTime()
	}
}

func Close() {
	globalLock.Lock()
	defer globalLock.Unlock()

	if faker != nil {
		faker.Close()
		faker = nil
	}
}
