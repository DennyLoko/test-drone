package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	foo := &Foo{}
	foo.SetBar("baz")

	assert.Equal(t, "baz", foo.Bar())
}
