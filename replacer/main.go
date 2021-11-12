package main

import (
	"os"
	"strings"
)

var (
	// map[oldString]newString
	dict = map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}

	newReplacer = func(dict map[string]string) *strings.Replacer {
		oldNew := make([]string, len(dict)*2)
		var i int
		for old, new := range dict {
			oldNew[i*2] = old
			oldNew[i*2+1] = new
			i++
		}
		return strings.NewReplacer(oldNew...)
	}
)

func main() {
	r := newReplacer(dict)
	dat, _ := os.ReadFile("in.txt")
	old := string(dat)
	new := r.Replace(old)

	out, _ := os.Create("out.txt")
	defer out.Close()
	out.WriteString(new)
}
