package piscine_go

func Strrev(s string) string{
	sc := []byte(s)
	for i, j := 0, len(sc)-1; i < j; i, j = i+1, j-1 {
		sc[i], sc[j] = sc[j], sc[i]
	}
	sf := string(sc)
	return sf
}
