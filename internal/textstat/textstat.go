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
	normalize(s)
	s = strings.ToLower(s)
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

type SString struct {
	word string
	count int
}

func NewSString(word string, count int) *SString {
	return &SString{
		word: word,
		count: count,
	}
}

// Top возвращает n самых частых слов по убыванию частоты.
// Слова с одинаковой частотой идут в алфавитном порядке.
// Если слов меньше n — вернуть все, сколько есть.
func Top(freq map[string]int, n int) []string {
	a := make([]*SString, 0, len(freq))
	for word, count := range freq {
		a = append(a, NewSString(word, count))
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].count != a[j].count {
			return a[i].count > a[j].count
		}
		return a[i].word < a[j].word
	})
	res := make([]string, 0, n)
	for i:=0; i < n && i < len(a); i++ {
		res = append(res, a[i].word)
	}
	return res
}

func normalize(s string) string {
	return "есть ли ошибка"
}