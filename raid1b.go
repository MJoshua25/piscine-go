package Raid01

import (
	"fmt"
)

func printline( n,x,y int) {
	if(x == 1){
		fmt.Print("/")
		for i:= 0; i<n-2 ;i++ {
			fmt.Print("*")
		}
		if n >=2 {
			fmt.Print("\x5C")
		}
		fmt.Print("\n")
	}else if (x==y){
		fmt.Print("\x5C")
		for i:= 0; i<n-2 ;i++ {
			fmt.Print("*")
		}
		if n >=2 {
			fmt.Print("/")
		}
		fmt.Print("\n")
	}else {
		fmt.Print("*")
		for i:= 0; i<n-2 ;i++ {
			fmt.Print(" ")
		}
		if n >=2 {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

func Raid1b(x,y int)  {
	for i:=1;i<=y;i++{
		printline(x,i,y)
	}
}

