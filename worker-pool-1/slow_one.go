package worker_pool_1

import (
	"math/rand/v2"
	"time"
)

type Request struct {
	Uuid  string // уникальный идентификатор задачи
	Input int64
}

type Response struct {
	Result int64
}

func slowLoad(requests []Request) map[string]Response {
	result := make(map[string]Response)
	for _, req := range requests {
		result[req.Uuid] = looooooad(req)
	}

	return result
}

func looooooad(r Request) Response {
	time.Sleep(rand.N(100 * time.Millisecond))
	return Response{Result: r.Input}
}
