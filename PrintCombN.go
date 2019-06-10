package piscine_go

import "fmt"

func printi(b int, n int, a int, i int)  {
	for j:= b+1; j<10-n+a; j++{
		fmt.Print(j)
		if(a+1!=n){
			a++
			printi(j,n,a,i)
		}else{
			if i<10-n{
				fmt.Print(", ")
			} else{
				fmt.Print("\n")
			}
		}
	}
}

func PrintCombN(n int)  {
	for i:=0; i<10-n; i++{
		fmt.Print(i)
		if n>1{
			printi(i,n,1,i)
		} else{
			if i<10-n{
				fmt.Print(", ")
			} else{
				fmt.Print("\n")
			}
		}
	}
}