package piscine_go

import "strconv"

func BasicAtoi(s string) int{
	a, err:= strconv.Atoi(s)
	return a
}
