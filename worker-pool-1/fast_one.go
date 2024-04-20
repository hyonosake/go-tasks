package worker_pool_1

import (
	"fmt"
	"sync"
)

func FastLoad(requests []Request) map[string]Response {
	/* тут твой кодик */
	results := make(map[string]Response)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, req := range requests {
		wg.Add(1)
		reqs := req

		go func() {
			defer wg.Done()
			mu.Lock()
			results[reqs.Uuid] = looooooad(reqs)
			mu.Unlock()
		}()

	}
	wg.Wait()
	fmt.Println(results)
	return results

}
