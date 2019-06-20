package piscine_go

func Swap(a *int, b *int)  {
	aux := *a
	*a = *b
	*b = aux
}