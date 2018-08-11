package cache_test

import (
	"testing"

	"github.com/jwowillo/hub/cache"
)

type MockCache struct {
	GetCalledWith       []string
	PutKeysCalledWith   []string
	PutValuesCalledWith []interface{}
	DeleteCalledWith    []string
	ClearCalls          int
}

func (c *MockCache) Get(k string) (interface{}, bool) {
	c.GetCalledWith = append(c.GetCalledWith, k)
	return nil, false
}

func (c *MockCache) Put(k string, v interface{}) {
	c.PutKeysCalledWith = append(c.PutKeysCalledWith, k)
	c.PutValuesCalledWith = append(c.PutValuesCalledWith, v)
}

func (c *MockCache) Delete(k string) {
	c.DeleteCalledWith = append(c.DeleteCalledWith, k)
}

func (c *MockCache) Clear() {
	c.ClearCalls++
}

// DecoratorTest is a test that makes sure Decorators created by
// DecoratorFactorys decorate Caches correctly.
//
// This only works with DecoratorFactorys that guarantee that each method in the
// Decorator only calls the respective decorated method and only calls each
// method once.
func DecoratorTest(t *testing.T, df cache.DecoratorFactory) {
	DecoratorGetTest(t, df)
	DecoratorPutTest(t, df)
	DecoratorDeleteTest(t, df)
	DecoratorClearTest(t, df)
}

// DecoratorGetTest is a test that makes sure Decorators created by
// DecoratorFactorys decorate Cache's Get correctly.
//
// This only works with DecoratorFactorys that guarantee that the Decorator's
// Get only calls the decorated Cache's Get once and calls no other methods.
func DecoratorGetTest(t *testing.T, df cache.DecoratorFactory) {
	mc := MockCache{}
	c := df(&mc)

	c.Get("k")
	if len(mc.GetCalledWith) != 1 || mc.GetCalledWith[0] != "k" {
		t.Errorf("mc.GetCalledWith = %v, want %v",
			mc.GetCalledWith, []string{"k"})
	}
}

// DecoratorPutTest is a test that makes sure Decorators created by
// DecoratorFactorys decorate Cache's Put correctly.
//
// This only works with DecoratorFactorys that guarantee that the Decorator's
// Put only calls the decorated Cache's Put once and calls no other methods.
func DecoratorPutTest(t *testing.T, df cache.DecoratorFactory) {
	mc := MockCache{}
	c := df(&mc)

	c.Put("k", 1)
	if len(mc.PutKeysCalledWith) != 1 || mc.PutKeysCalledWith[0] != "k" {
		t.Errorf("mc.PutKeysCalledWith = %v, want %v",
			mc.PutKeysCalledWith, []string{"k"})
	}
	if len(mc.PutValuesCalledWith) != 1 || mc.PutValuesCalledWith[0] != 1 {
		t.Errorf("mc.PutValuesCalledWith = %v, want %v",
			mc.PutValuesCalledWith, []interface{}{1})
	}
}

// DecoratorDeleteTest is a test that makes sure Decorators created by
// DecoratorFactorys decorate Cache's Delete correctly.
//
// This only works with DecoratorFactorys that guarantee that the Decorator's
// Get only calls the decorated Cache's Delete once and calls no other methods.
func DecoratorDeleteTest(t *testing.T, df cache.DecoratorFactory) {
	mc := MockCache{}
	c := df(&mc)

	c.Delete("k")
	if len(mc.DeleteCalledWith) != 1 || mc.DeleteCalledWith[0] != "k" {
		t.Errorf("mc.DeleteCalledWith = %v, want %v",
			mc.DeleteCalledWith, []string{"k"})
	}
}

// DecoratorClearTest is a test that makes sure Decorators created by
// DecoratorFactorys decorate Cache's Clear correctly.
//
// This only works with DecoratorFactorys that guarantee that the Decorator's
// Get only calls the decorated Cache's Clear once and calls no other methods.
func DecoratorClearTest(t *testing.T, df cache.DecoratorFactory) {
	mc := MockCache{}
	c := df(&mc)

	c.Clear()
	if mc.ClearCalls != 1 {
		t.Errorf("mc.ClearCalls = %d, want %d",
			mc.ClearCalls, 0)
	}
}
