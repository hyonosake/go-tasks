package worker_pool_1

import (
	"math/rand/v2"
	"time"
)

type Request struct {
	uuid  string // уникальный идентификатор задачи
	input int64
}

type Response struct {
	result int64
}

func slowLoad(requests []Request) map[string]Response {
	result := make(map[string]Response)
	for _, req := range requests {
		result[req.uuid] = looooooad(req)
	}

	return result
}

func looooooad(r Request) Response {
	time.Sleep(rand.N(100 * time.Millisecond))
	return Response{result: r.input}
}
