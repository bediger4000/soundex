package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	zork := soundex(os.Args[1])
	fmt.Printf("%q\n", zork)
}
func soundex(orig string) string {
	runes := []rune(orig)
	var sndx []rune

	for i := range runes {
		runes[i] = unicode.ToLower(runes[i])
	}

	sndx = append(sndx, runes[0])

	runes = removeVowels(runes[1:])

	runes = removeConsecutive(runes)

	numberCount := 0
	for i := range runes {
		var r rune
		switch runes[i] {
		case 'b', 'f', 'p', 'v':
			r = '1'
		case 'c', 'g', 'j', 'k', 'q', 's', 'x', 'z':
			r = '2'
		case 'd', 't':
			r = '3'
		case 'l':
			r = '4'
		case 'm', 'n':
			r = '5'
		case 'r':
			r = '6'
		}
		sndx = append(sndx, r)
		numberCount++
	}
	for numberCount < 3 {
		sndx = append(sndx, '0')
		numberCount++
	}

	if len(sndx) > 4 {
		sndx = sndx[:4]
	}

	return string(sndx)
}

var soundAlikes map[rune][]rune = map[rune][]rune{
	'c': []rune{'k', 's', 'q'},
	'k': []rune{'c', 'q'},
	'q': []rune{'c', 'k'},
	'g': []rune{'j'},
	'j': []rune{'g'},
}

// Remove consecutive consonants with the same sound (for
// example, change ck -> c).
func removeConsecutive(runes []rune) []rune {
	removed := make([]rune, len(runes))
	removed[0] = runes[0]
	rIdx := 1

	last := runes[0]
	for i := 1; i < len(runes); i++ {
		if last == runes[i] {
			continue
		}
		sa := soundAlikes[last]
		foundit := false
		for _, r := range sa {
			if runes[i] == r {
				foundit = true
				break
			}
		}
		if foundit {
			continue
		}
		last = runes[i]
		removed[rIdx] = runes[i]
		rIdx++
	}

	return removed[:rIdx]
}

func removeVowels(runes []rune) []rune {
	removed := make([]rune, len(runes))
	rIdx := 0
	for _, r := range runes {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'y', 'w', 'h':
			continue
		default:
			removed[rIdx] = r
			rIdx++
		}
	}
	return removed[:rIdx]
}
