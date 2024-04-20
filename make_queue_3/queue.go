package make_queue_3

/*
	Ну и последняя итерация - добавим ограничения:
		- одновременно в очереди не может быть больше queueLimit задач
		- если очередь заполнена, AddTask должен вернуть ошибку (а не заблокироваться, ожидая пока появится место)
		- одновременно может работать максимум poolSize воркеров, которые берут задачи из очереди и выполняют их
		- вызов Shudown() должен сделать всё то же, что и в прошлый раз, но также "завершить" воркеров, чтобы они не работали
		- вхолостую, так как это будет считай утечка памяти

		p.s. это задача middle+ уровня с собеса Cloud.ru :)
*/

type Task func()

func NewTask(f func()) Task {
	return Task(f)
}

type queue struct {
}

func NewQueue(queueLimit, poolSize int) *queue {
	return &queue{}
}

func (q *queue) AddTask(t Task) error {
	return nil
}

func (q *queue) Shutdown() {

}
