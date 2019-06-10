package piscine_go

import "fmt"
import "strconv"

func PrintComb()  {
	for i :=0; i<10; i++{
		for j :=0; j<10; j++{
			for k :=0; k<10; k++{
				if(i!=j && i!= k && j!=k){
					fmt.Print(strconv.Itoa(i)+strconv.Itoa(j)+strconv.Itoa(k)+", ")
				}
			}
		}
	}
}
