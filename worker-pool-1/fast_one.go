package worker_pool_1

func fastLoad(requests []Request) map[string]Response {
	/* тут твой кодик */
	results := make(map[string]Response)
	for _, req := range requests {
		reqs := req
		go func() {
			results[reqs.uuid] = looooooad(reqs)
		}()
	}
	return nil

}
