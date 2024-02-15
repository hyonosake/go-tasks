package worker_pool_1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlowOne_WithSmallBatch_Positive(t *testing.T) {
	requests := []Request{
		{
			uuid:  "uuid_1",
			input: 1,
		},
		{
			uuid:  "uuid_2",
			input: 2,
		},
		{
			uuid:  "uuid_3",
			input: 3,
		},
		{
			uuid:  "uuid_4",
			input: 4,
		},
	}

	expectedResult := map[string]Response{
		"uuid_1": {
			result: 1,
		},
		"uuid_2": {
			result: 2,
		},
		"uuid_3": {
			result: 3,
		},
		"uuid_4": {
			result: 4,
		},
	}
	result := slowLoad(requests)
	assert.Equal(t, expectedResult, result)
}

func TestSlowOne_WithLargeBatch_Positive(t *testing.T) {
	amount := 100

	requests := make([]Request, 0, amount)
	expectedResult := make(map[string]Response)

	for i := 0; i < amount; i++ {
		uuid := fmt.Sprintf("uuid_%d", i)
		requests = append(requests, Request{uuid: uuid, input: int64(i)})
		expectedResult[uuid] = Response{result: int64(i)}
	}

	result := slowLoad(requests)
	assert.Equal(t, expectedResult, result)
}
