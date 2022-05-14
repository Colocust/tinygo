package tinygo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoinPaths(t *testing.T) {
	absolutePath := "/"
	relativePath := "user/"

	finalPath := joinPaths(absolutePath, relativePath)

	assert.Equal(t, "/user/", finalPath)
}
