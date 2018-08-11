package cache_test

import (
	"testing"
	"time"

	"github.com/jwowillo/hub/cache"
)

type MockTimeSource struct {
	current time.Time
}

func (ts *MockTimeSource) Step() {
	ts.current = ts.current.Add(1)
}

func (ts *MockTimeSource) Time() time.Time {
	return ts.current
}

// TestTimeDecoratorGet tests that only the decoratored Cache's Get is called if
// the time isn't after the time.Duration and only the decorated Cache's Clear
// is called if the time is after the time.Duration.
func TestTimeDecoratorGet(t *testing.T) {
	ts := &MockTimeSource{}
	mc := &MockCache{}
	c := cache.NewTimeDecoratorFactory(ts.Time, 0)(mc)
	c.Get("k1")
	if len(mc.GetCalledWith) != 1 || mc.GetCalledWith[0] != "k1" {
		t.Errorf("mc.GetCalledWith = %v, want %v",
			mc.GetCalledWith, []string{"k1"})
	}
	if mc.ClearCalls != 0 {
		t.Errorf("mc.ClearCalls = %d, want %d", mc.ClearCalls, 0)
	}

	ts.Step()

	c.Get("k2")
	if len(mc.GetCalledWith) != 1 || mc.GetCalledWith[0] != "k1" {
		t.Errorf("mc.GetCalledWith = %v, want %v",
			mc.GetCalledWith, []string{"k1"})
	}
	if mc.ClearCalls != 1 {
		t.Errorf("mc.ClearCalls = %d, want %d", mc.ClearCalls, 1)
	}
}

// TestTimeDecoratorPut tests that TimeDecorator decorates the decorated Cache's
// Put properly.
func TestTimeDecoratorPut(t *testing.T) {
	mc := MockTimeSource{}
	f := cache.NewTimeDecoratorFactory(mc.Time, 0)
	DecoratorPutTest(t, f)
}

// TestTimeDecoratorPut tests that TimeDecorator decorates the decorated Cache's
// Delete properly.
func TestTimeDecoratorDelete(t *testing.T) {
	mc := MockTimeSource{}
	f := cache.NewTimeDecoratorFactory(mc.Time, 0)
	DecoratorDeleteTest(t, f)
}

// TestTimeDecoratorPut tests that TimeDecorator decorates the decorated Cache's
// Clear properly.
func TestTimeDecoratorClear(t *testing.T) {
	mc := MockTimeSource{}
	f := cache.NewTimeDecoratorFactory(mc.Time, 0)
	DecoratorClearTest(t, f)
}
