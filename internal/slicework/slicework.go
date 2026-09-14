package slicework
// Dedup убирает ПОДРЯД ИДУЩИЕ дубликаты, изменяя срез на месте,
// без выделения нового массива. Возвращает срез с нужной длиной.
func Dedup(s []int) []int {
	for i, _ := range s {
		if i > 0 && s[i] == s[i - 1] {
			s = append(s[:i-1], s[i:]...)
		}
	}
	return s
}

// Chunk разбивает срез на куски длиной n (последний может быть короче).
// append в любой кусок НЕ должен портить соседние куски.
func Chunk(s []int, n int) [][]int {
	result := [][]int{}
	var i int = 0
	for i = n; i < n; i += n {
		result = append(result, s[i-n:i:n])
	}
	if len(s[i-n:]) != 0 {
		result = append(result, s[i-n:])
	}
	return result
}

// Invert разворачивает мапу: значение становится ключом,
// а списком идут все ключи с этим значением.
// Списки должны быть отсортированы — результат обязан быть детерминированным.
func Invert(m map[string]int) map[int][]string {
	insert := make(map[int][]string)
	for key, val := range m {
		for v := range val {
			insert[v] = append(insert[v], key)
		}
	}
	return insert
}