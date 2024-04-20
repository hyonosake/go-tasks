package worker_pool_1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlowOne_WithSmallBatch_Positive(t *testing.T) {
	requests := []Request{
		{
			Uuid:  "uuid_1",
			Input: 1,
		},
		{
			Uuid:  "uuid_2",
			Input: 2,
		},
		{
			Uuid:  "uuid_3",
			Input: 3,
		},
		{
			Uuid:  "uuid_4",
			Input: 4,
		},
	}

	expectedResult := map[string]Response{
		"Uuid_1": {
			Result: 1,
		},
		"Uuid_2": {
			Result: 2,
		},
		"Uuid_3": {
			Result: 3,
		},
		"Uuid_4": {
			Result: 4,
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
		uuid := fmt.Sprintf("Uuid_%d", i)
		requests = append(requests, Request{Uuid: uuid, Input: int64(i)})
		expectedResult[uuid] = Response{Result: int64(i)}
	}

	result := slowLoad(requests)
	assert.Equal(t, expectedResult, result)
}
