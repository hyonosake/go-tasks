package make_ticker

import (
	"fmt"
	"time"
)

/*
	Нужно написать функцию-обёртку OnTick(), которая будет принимать в себя какую-то функцию и time.Duration
	Например, если я вызову функцию Tick(printTime, time.Second), то у меня каждую секунду должен будет появляться лог
	в консоли вида `it's (время)`
	Юнит-тесты писать необязательно, но нужен мейник с примером, как это будет работать
	Ну или юнит, если тебе так проще
*/

func printTime() {
	fmt.Printf("it's %v\n", time.Now().Format(time.TimeOnly))
}

func Tick(callable func(), d time.Duration) {
	/* тут твой кодик */
}
