/* Lesson_19
Parametre olarak fonksiyon kullanmak
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	colors := []string{
		"red", "blue", "green", "orange", "black", "yellow",
		"gray", "brown", "silver", "pink", "gold", "dark red",
	}
	result := Select(colors, func(c string) bool {
		return len(c) >= 5
	})
	for _, r := range result {
		fmt.Printf("%s,", r)
	}
	fmt.Println()
	g := func(w string) bool {
		return strings.HasPrefix(w, "g")
	}
	result = Select(colors, g)
	for _, r := range result {
		fmt.Printf("%s,", r)
	}
	fmt.Println()
}

type predicate func(w string) bool

func Select(words []string, f predicate) []string {
	findings := []string{}
	for _, word := range words {
		if f(word) {
			findings = append(findings, word)
		}
	}
	return findings
}
