package make_queue

import "sync"

/*
	Продолжаем писать либы, хэх
	Требуется реализовать "очередь задач", которая принимает на вход какую-то функцию и кладёт её в очередь
	Очередь должна быть in-memory
	Задачи, попавшие в очередь, должны по FIFO браться оттуда воркером (одним) и выполняться
	вызов AddTask должен быть НЕ БЛОКИРУЮЩИМ, т.е. если я загрузил в AddTask таску на долгое время
	(например она выполняется 2 секунды), то я могу продолжать свою работу дальше, пока на фоне будет выполняться моя таска

	Юнит-тесты / мейник, который показывает что всё работает
*/

type Task func()

func NewTask(f func()) Task {
	return Task(f)
}

type Queue struct {
	QueueTasks chan Task
}

func NewQueue(len int) *Queue {
	return &Queue{
		QueueTasks: make(chan Task, len),
	}
}

func (q *Queue) AddTask(t Task) {
	q.QueueTasks <- t
}

type Worker struct {
	Wg *sync.WaitGroup
}

func (q *Worker) WorkerFunc(t <-chan Task) {
	for task := range t {
		q.Wg.Add(1)
		task()
		q.Wg.Done()
	}
}
