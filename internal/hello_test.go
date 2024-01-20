package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	assert := assert.New(t)
	t.Run(`should say hello`, func(t *testing.T) {
		assert.Equal(`Hello`, `Hello`)
	})
}
