package textstat

import (
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// CountRunes возвращает количество символов (не байтов) в строке.
func CountRunes(s string) int {
	return utf8.RuneCountInString(s)
}

// Words разбивает строку на слова. Словом считается последовательность
// букв и цифр; всё остальное — разделители. Регистр приводится к нижнему.
func Words(s string) []string {
	s = normalize(s)
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	return words
}

// Frequency возвращает, сколько раз каждое слово встретилось.
func Frequency(words []string) map[string]int {
	a := map[string]int{}
	for _, word := range words {
		a[word]++
	}
	return a
}

type wordCount struct {
	word  string
	count int
}

// Top возвращает n самых частых слов по убыванию частоты.
// Слова с одинаковой частотой идут в алфавитном порядке.
// Если слов меньше n — вернуть все, сколько есть.
func Top(freq map[string]int, n int) []string {
	if n < 0 {
		return []string{}
	}
	wordCnt := make([]wordCount, 0, len(freq))
	for word, count := range freq {
		wordCnt = append(wordCnt, wordCount{word, count})
	}
	sort.Slice(wordCnt, func(i, j int) bool {
		if wordCnt[i].count != wordCnt[j].count {
			return wordCnt[i].count > wordCnt[j].count
		}
		return wordCnt[i].word < wordCnt[j].word
	})
	result := make([]string, 0)
	for i := 0; i < n && i < len(wordCnt); i++ {
		result = append(result, wordCnt[i].word)
	}
	return result
}

func normalize(s string) string {
	return strings.ToLower(s)
}
