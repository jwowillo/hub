package cache

import (
	"io"
	"os"
	"sync"
	"time"
)

// DefaultWriter returns os.Stdout.
func DefaultWriter() io.Writer {
	return os.Stdout
}

// DefaultTimeSource returns time.Now.
func DefaultTimeSource() TimeSource {
	return time.Now
}

// DefaultHasBeenModified returns true if the file at the path has been modified
// since its associated value was stored.
//
// Returns true if the file can't be accessed also. This behavior is to force
// the value to be refetched in case of errors.
func DefaultHasBeenModified() HasBeenModified {
	return func(path string, last time.Time) bool {
		f, err := os.Stat(path)
		if err != nil {
			return true
		}
		return f.ModTime().After(last)
	}
}

// DefaultLockers returns a sync.RWMutex's write sync.Locker and read
// sync.Locker.
func DefaultLockers() (sync.Locker, sync.Locker) {
	m := &sync.RWMutex{}
	return m, m.RLocker()
}
