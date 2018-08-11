package cache_test

import (
	"github.com/jwowillo/hub/cache"

	"testing"
)

// TestGet tests that Get tries to get from the Cache first, then calls the
// Fallback, stores values from the Fallback back into the Cache, and returns
// the correct values.
func TestGet(t *testing.T) {
	var fallbackCalledWith string
	mc := &MockCache{}
	v := cache.Get(mc, "k", func(k string) interface{} {
		fallbackCalledWith = k
		return 1
	})
	if len(mc.GetCalledWith) != 1 || mc.GetCalledWith[0] != "k" {
		t.Errorf("mc.GetCalledWith = %v, want %v",
			mc.GetCalledWith, []string{"k"})
	}
	if fallbackCalledWith != "k" {
		t.Errorf("fallbackCalledWith = %s, want %s",
			fallbackCalledWith, "k")
	}
	if len(mc.PutKeysCalledWith) != 1 || mc.PutKeysCalledWith[0] != "k" {
		t.Errorf("mc.PutKeysCalledWith = %v, want %v",
			mc.PutKeysCalledWith, []string{"k"})
	}
	if len(mc.PutValuesCalledWith) != 1 || mc.PutValuesCalledWith[0] != 1 {
		t.Errorf("mc.PutValuesCalledWith = %v, want %v",
			mc.PutValuesCalledWith, []interface{}{1})
	}
	if len(mc.DeleteCalledWith) != 0 {
		t.Errorf("mc.DeleteCalledWith = %v, want %v",
			mc.DeleteCalledWith, nil)
	}
	if mc.ClearCalls != 0 {
		t.Errorf("mc.ClearCalls = %d, want %d",
			mc.ClearCalls, 0)
	}
	if v != 1 {
		t.Errorf("v = %v, want %v", v, 1)
	}
}
