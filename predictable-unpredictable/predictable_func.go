package predictable_unpredictable

import (
	"math/rand/v2"
	"time"
)

const timeoutMs = time.Millisecond * 1000 // nolint

type Result int64

func NewResult(n int64) Result {
	return Result(n)
}

func unpredictableFunc() Result { // nolint
	// тут какие-то оооч сложные вычисления
	time.Sleep(rand.N(2000 * time.Millisecond))
	return NewResult(rand.N(int64(2000)))
}

// predictableFunc возвращает result, если результат получился
// error, если мы не успели обработать ответ за отведённое время (timeoutMs)
func predictableFunc( /* тут пииши что хочешь */ ) (Result, error) {
	// ...
	return 0, nil
}
