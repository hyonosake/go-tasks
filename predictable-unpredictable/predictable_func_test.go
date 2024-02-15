package predictable_unpredictable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unpredictableFunc(t *testing.T) {
	_, err := predictableFunc()
	assert.NoError(t, err)
}
