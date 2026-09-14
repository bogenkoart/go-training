package slicework
// Dedup убирает ПОДРЯД ИДУЩИЕ дубликаты, изменяя срез на месте,
// без выделения нового массива. Возвращает срез с нужной длиной.
func Dedup(s []int) []int {
	if len(s) == 0 {
		return []int{}
	}
	dedup := s[:0]
	slow, fast := 1, 1
	for fast < len(s) {
		if dedup[fast] != dedup[fast - 1] {
			dedup[slow] = dedup[fast]
			slow++
		}
		fast++
	}
	return dedup
}

// Chunk разбивает срез на куски длиной n (последний может быть короче).
// append в любой кусок НЕ должен портить соседние куски.
func Chunk(s []int, n int) [][]int {
	result := [][]int{}
	var i int = 0
	for i = n; i < len(s); i += n {
		result = append(result, s[i-n:i:i])
	}
	if len(s[i-n:]) != 0 {
		result = append(result, s[i-n:len(s):len(s)])
	}
	return result
}

// Invert разворачивает мапу: значение становится ключом,
// а списком идут все ключи с этим значением.
// Списки должны быть отсортированы — результат обязан быть детерминированным.
func Invert(m map[string]int) map[int][]string {
	insert := make(map[int][]string)
	for key, val := range m {
		insert[val] = append(insert[val], key)
	}
	return insert
}