package make_ticker_3

import "time"

/*
	Задача усложняется ещё немного: я хочу сделать что-то типа библиотечного API
	предлагая Ticker в качестве использования другими пакетами
	И теперь тикер должен запускать неограниченное количество функций, а при вызове Shutdown()
	они все должны перестать "тикать"

	p.s. тут уже точно нужна будет структура
*/

type tickerScheduler struct {
	/* что тебе тут нужно? */
}

func NewTickScheduler() *tickerScheduler {
	return &tickerScheduler{}
}

func (t *tickerScheduler) AddTask(func(), time.Duration) {
	/* добавить новый тикер */
}

func (t *tickerScheduler) Shutdown() {
	/* остановить все запущенные тикеры */
}
