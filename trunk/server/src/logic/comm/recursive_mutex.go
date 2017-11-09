package comm

import (
	"runtime"
)

type RecursiveMutex struct {
	stackId int32
}

func (this *RecursiveMutex) TryLock() {

	runtime.LockOSThread()
}

func (this *RecursiveMutex) Unlock() {
	runtime.UnlockOSThread()
}
