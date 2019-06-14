package piscine_go

import "strconv"

func BasicAtoi(s string) int{
	a, _:= strconv.Atoi(s)
	return a
}
