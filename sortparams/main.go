package main

import (
	"fmt"
	"os"
	"sort"
)

func main()  {
	aux := os.Args[0:len(os.Args)]
	fmt.Println(aux)
	sort.Strings(aux)
	for i:=1; i<len(aux);i++{
		fmt.Println(aux[i])
	}
}