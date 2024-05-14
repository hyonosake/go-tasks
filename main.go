package main

import (
	"fmt"
	"github.com/hyonosake/go-tasks/make_queue"
	"sync"
	"time"
)

func lol() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hello aaa, ")
}

func main() {
	//amount := 100
	//
	//requests := make([]workerpool1.Request, 0, amount)
	//expectedResult := make(map[string]workerpool1.Response)
	//
	//for i := 0; i < amount; i++ {
	//	uuid := fmt.Sprintf("uuid_%d", i)
	//	requests = append(requests, workerpool1.Request{Uuid: uuid, Input: int64(i)})
	//	expectedResult[uuid] = workerpool1.Response{Result: int64(i)}
	//}
	//
	//_ = workerpool1.FastLoad(requests)

	len := 3
	wg := sync.WaitGroup{}
	q := make_queue.NewQueue(len)
	var w make_queue.Worker

	w.Wg = &wg

	q.QueueTasks = make(chan make_queue.Task, len)
	t := time.Now()
	q.AddTask(make_queue.NewTask(lol))
	lol()

	go w.WorkerFunc(q.QueueTasks)

	w.Wg.Wait()
	fmt.Println(time.Since(t))
}
