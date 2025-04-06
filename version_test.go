package ptt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	assert.IsType(t, 0, Version().Int())
	assert.IsType(t, "", Version().String())
}
