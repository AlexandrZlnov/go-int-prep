// 13. Написать юнит-тесты для метода NewObject.

package main

import (
	"testing"
)

func TestNewObject(t *testing.T) {
	obj := NewObject("base", "Parent1", "")
	b, ok := obj.(*Base)
	if !ok {
		t.Errorf("Expected Base, Got %T", obj)
	}
	if b.name != "Parent1" {
		t.Errorf("Expected Parent1, Got %s", b.name)
	}

	obj = NewObject("child", "Child1", "Inherited1")
	c, ok := obj.(*Child)
	if !ok {
		t.Errorf("Expected Child, Got %T", obj)
	}
	if c.name != "Child1" {
		t.Errorf("Expected Child1, Got %s", c.name)
	}
	if c.lastName != "Inherited1" {
		t.Errorf("Expected Inherited1, Got %s", c.lastName)
	}
}
