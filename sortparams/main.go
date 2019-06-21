package main

import (
	"fmt"
	"os"
	"sort"
)

func main()  {
	aux := os.Args[1:len(os.Args)]
	sort.Strings(aux)
	for i:=1; i<len(aux);i++{
		fmt.Println(aux[i])
	}
}