package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocalIPv4Address(t *testing.T) {
	addr, err := GetLocalIPv4Address()
	assert.Nil(t, err)
	println(addr)
}
