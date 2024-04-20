package main

import (
	"fmt"
	workerpool1 "github.com/hyonosake/go-tasks/worker-pool-1"
)

func main() {
	amount := 100

	requests := make([]workerpool1.Request, 0, amount)
	expectedResult := make(map[string]workerpool1.Response)

	for i := 0; i < amount; i++ {
		uuid := fmt.Sprintf("uuid_%d", i)
		requests = append(requests, workerpool1.Request{Uuid: uuid, Input: int64(i)})
		expectedResult[uuid] = workerpool1.Response{Result: int64(i)}
	}

	_ = workerpool1.FastLoad(requests)
}
