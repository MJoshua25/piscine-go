package piscine_go

import "strings"

func FirstRune(s string) rune{
	r := []rune(s)
	return r[0]
}

func NRune(s string, n int) rune{
	r := []rune(s)
	return r[n]
}

func LastRune(s string) rune{
	r := []rune(s)
	return r[len(r)-1]
}

func Index(s, toFind string) int {
	return strings.Index(s,toFind)
}

func Compare (a,b string) int{
	return strings.Compare(a,b)
}

