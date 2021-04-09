package main

import (
	"encoding/json"
	"io"
)

type foxyPattern struct {
	Title     string `json:"title"`
	Pattern   string `json:"pattern"`
	Active    bool   `json:"active"`
	Enabled   bool   `json:"enabled"`
	Type      int    `json:"type"`
	Protocols int    `json:"protocols"`
}

type foxyStruct struct {
	WhitePatterns []foxyPattern `json:"whitePatterns"`
	BlackPatterns []foxyPattern `json:"blackPatterns"`
}

func nameToFoxyPattern(in string) string {
	return "*" + in
}

func foxyProxy(target io.Writer, domains ...string) error {
	all := foxyStruct{
		WhitePatterns: make([]foxyPattern, len(domains)),
	}

	for i := range domains {
		all.WhitePatterns[i] = foxyPattern{
			domains[i],
			nameToFoxyPattern(domains[i]),
			true,
			true,
			1,
			1,
		}
	}
	enc := json.NewEncoder(target)
	return enc.Encode(all)
}
