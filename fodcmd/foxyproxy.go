package main

import (
	"encoding/json"
	"io"
	"strings"
)

type foxyPattern struct {
	Enabled       bool   `json:"enabled"`
	Name          string `json:"name"`
	Pattern       string `json:"pattern"`
	IsRegEx       bool   `json:"isRegEx"`
	CaseSensitive bool   `json:"caseSensitive"`
	BlackList     bool   `json:"blackList"`
	MultiLine     bool   `json:"multiLine"`
}

type foxyStruct struct {
	Patterns []foxyPattern `json:"patterns"`
}

func nameToFoxyPattern(in string) string {
	return strings.Replace(in, ".", "*", 1) + "/*"
}

func foxyProxy(target io.Writer, domains ...string) error {
	all := foxyStruct{
		Patterns: make([]foxyPattern, len(domains)),
	}

	for i := range domains {
		all.Patterns[i] = foxyPattern{
			true,
			domains[i],
			nameToFoxyPattern(domains[i]),
			false,
			false,
			false,
			false,
		}
	}
	enc := json.NewEncoder(target)
	return enc.Encode(all)
}
