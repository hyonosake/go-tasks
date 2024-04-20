package sync_map

/*
	Требуется написать свою реализацию синк-мапы :)
	Я для простоты сделал мапу interface / interface, можешь бонусом попробовать написать её
	на дженериках, чтобы можно было например сделать вот так NewSyncMap[int, string](),
	и это бы значило что ты создал map[int]string
*/

type syncMap struct {
}

func NewSyncMap() *syncMap {
	return &syncMap{}
}

// Get возвращает значение по ключу, и нашлось ли оно
func (s *syncMap) Get(key interface{}) (interface{}, bool) {
	return 0, false
}

// Set добавляет новое значение
func (s *syncMap) Set(key, value interface{}) {
}

// IsEmpty помогает понять, пустая ли у нас мапа
func (s *syncMap) IsEmpty() bool {
	return true
}

// Size возвращает кол-во элементов в мапе
func (s *syncMap) Size() bool {
	return true
}
