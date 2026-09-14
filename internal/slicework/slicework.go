package slicework

import "sort"

// Dedup убирает ПОДРЯД ИДУЩИЕ дубликаты, изменяя срез на месте,
// без выделения нового массива. Возвращает срез с нужной длиной.
func Dedup(s []int) []int {
	if len(s) == 0 {
		return s
	}
	dedup := s[:1]
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			dedup = append(dedup, s[i])
		}
	}
	return dedup
}

// Chunk разбивает срез на куски длиной n (последний может быть короче).
// append в любой кусок НЕ должен портить соседние куски.
func Chunk(s []int, n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
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
	for _, val := range insert {
		sort.Strings(val)
	}
	return insert
}
