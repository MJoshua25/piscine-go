package piscine_go

import "strconv"

func BasicAtoi(s string) int{
	a, _:= strconv.Atoi(s)
	return a
}

func BasicAtoi2(s string) int{
	return BasicAtoi(s)
}

func Atoi(s string) int{
	return BasicAtoi(s)
}
