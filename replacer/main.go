package main

import (
	"os"
	"strings"
)

var (
	// []map[oldString]newString
	dict = []map[string]string{
		{
			"I":    "U",
			"love": "LOVE",
		},
		{
			"I":    "HE",
			"love": "LOVES",
		},
		{
			"I":    "SHE",
			"love": "LOVES",
		},
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
	dat, _ := os.ReadFile("in.txt")
	old := string(dat)

	out, _ := os.Create("out.txt")
	defer out.Close()

	for _, d := range dict {
		r := newReplacer(d)
		new := r.Replace(old)
		out.WriteString(new)
		out.WriteString("\n\n")
	}

}
