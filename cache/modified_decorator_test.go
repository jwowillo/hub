package cache_test

import (
	"testing"
	"time"

	"github.com/jwowillo/hub/cache"
)

// TestModifiedDecoratorGet tests that ModifiedDecorator's Get deletes stale
// entries without calling the decorated Cache's Get and returns valid entries
// with calling the decorated Cache's Get.
//
// Depends on Put working.
func TestModifiedDecoratorGet(t *testing.T) {
	var keyCalledWith string
	var timeCalledWith time.Time
	isModified := false
	ts := &MockTimeSource{}
	ts.Step()
	mc := &MockCache{}
	c := cache.NewModifiedDecoratorFactory(
		ts.Time,
		func(k string, t time.Time) bool {
			keyCalledWith = k
			timeCalledWith = t
			return isModified
		})(mc)

	if _, ok := c.Get("k"); ok {
		t.Errorf("c.Get(%s) = true, want false", "k")
	}

	c.Put("k", 1)

	c.Get("k")
	if keyCalledWith != "k" {
		t.Errorf("keyCalledWith = %s, want %s", keyCalledWith, "k")
	}
	if !timeCalledWith.Equal(time.Time{}.Add(1)) {
		t.Errorf("timeCalledWith = %v, want %v",
			timeCalledWith, time.Time{}.Add(1))
	}
	if len(mc.GetCalledWith) != 1 || mc.GetCalledWith[0] != "k" {
		t.Errorf("mc.GetCalledWith = %v, want %v",
			mc.GetCalledWith, []string{"k"})
	}
	if len(mc.DeleteCalledWith) != 0 {
		t.Errorf("mc.DeleteCalledWith = %v, want %v",
			mc.DeleteCalledWith, []string{})
	}

	isModified = true

	c.Get("k")
	if len(mc.GetCalledWith) != 1 {
		t.Errorf("mc.GetCalledWith = %v, want %v",
			mc.GetCalledWith, []string{"k"})
	}
	if len(mc.DeleteCalledWith) != 1 || mc.DeleteCalledWith[0] != "k" {
		t.Errorf("mc.DeleteCalledWith = %v, want %v",
			mc.DeleteCalledWith, []string{"k"})
	}
}

// TestModifiedDecoratorPut tests that ModifiedDecorator decorates the decorated
// Cache's Put properly.
func TestModifiedDecoratorPut(t *testing.T) {
	ts := &MockTimeSource{}
	f := cache.NewModifiedDecoratorFactory(
		ts.Time,
		func(string, time.Time) bool { return false })
	DecoratorPutTest(t, f)
}

// TestModifiedDecoratorPut tests that ModifiedDecorator decorates the decorated
// Cache's Delete properly.
func TestModifiedDecoratorDelete(t *testing.T) {
	ts := &MockTimeSource{}
	f := cache.NewModifiedDecoratorFactory(
		ts.Time,
		func(string, time.Time) bool { return false })
	DecoratorDeleteTest(t, f)
}

// TestModifiedDecoratorPut tests that ModifiedDecorator decorates the decorated
// Cache's Clear properly.
func TestModifiedDecoratorClear(t *testing.T) {
	ts := &MockTimeSource{}
	f := cache.NewModifiedDecoratorFactory(
		ts.Time,
		func(string, time.Time) bool { return false })
	DecoratorClearTest(t, f)
}
