package worker_pool_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastOne(t *testing.T) {
	requests := []Request{
		{
			uuid:  "uuid_1",
			input: 1,
		},
	}
	expectedResult := map[string]Response{
		"uuid_1": {
			result: 1,
		},
	}
	result := fastLoad(requests)

	assert.NotEqual(t, expectedResult, result)
}
