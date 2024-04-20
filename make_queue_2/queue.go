package make_queue_2

/*
	Супер, очередь сделали.
	Теперь необходимо добавить структуру метод Shutdown()
	Что он делает: он запрещает класть в очередь новые таски
	(то есть теперь AddTask должен возвращать ошибку, если до этого был вызван Shutown())
	Также после вызова Shutdown() мы должны запроцессить все задачи, которые в данный момент находятся у нас в очереди на исполнение
*/

type Task func()

func NewTask(f func()) Task {
	return Task(f)
}

type queue struct {
}

func NewQueue() *queue {
	return &queue{}
}

func (q *queue) AddTask(t Task) error {
	return nil
}

func (q *queue) Shutdown() {

}
