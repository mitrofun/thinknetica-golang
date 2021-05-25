package index

import (
	"strings"
)

type Index struct {
	words map[string][] int
}

func (i *Index) New() {
	i.words = make(map[string][]int)
}

func (i *Index) Add(s string, num int) {
	words := strings.Fields(strings.ToLower(s))
	for _, v := range words {
		i.words[v] = append(i.words[v], num)
	}
}

func (i *Index) Search (s string) []int {
	return i.words[strings.ToLower(s)]
}