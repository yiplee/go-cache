package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := New[string, string]()

	a, found := c.Get("a")
	if found || a != "" {
		t.Errorf("Get a found value should not exist: %v", a)
	}

	c.Set("a", "A")
	a, found = c.Get("a")
	if !found || a != "A" {
		t.Errorf("Get a found value should exist: %v", a)
	}

	c.Set("a", "B")
	a, found = c.Get("a")
	if !found || a != "B" {
		t.Errorf("Get a found value should exist: %v", a)
	}

	c.Set("b", "B", WithTTL(20*time.Millisecond))
	b, found := c.Get("b")
	if !found || b != "B" {
		t.Errorf("Get b found value should exist: %v", b)
	}

	time.Sleep(50 * time.Millisecond)
	b, found = c.Get("b")
	if found || b != "" {
		t.Errorf("Get b found value should not expired: %v", b)
	}

	if !c.Contain("a") {
		t.Errorf("Contain a should be true")
	}

	if expect, got := 1, len(c.Keys()); expect != got {
		t.Errorf("Keys should be %v, got %v", expect, got)
	}

	c.Delete("a")
	a, found = c.Get("a")
	if found || a != "" {
		t.Errorf("Get a found value should not exist: %v", a)
	}

	if c.Contain("a") {
		t.Errorf("Contain a should not exist: %v", a)
	}

	if expect, got := 0, len(c.Keys()); expect != got {
		t.Errorf("Keys should be %v, got %v", expect, got)
	}
}
