package l

import (
	"time"
)

type Lock interface {
	Lock()
	Unlock()
	TryLock(timeout time.Duration) bool
	isLock() bool
}
