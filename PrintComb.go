package piscine_go

import "fmt"
import "strconv"

func PrintComb()  {
	for i :=0; i<10; i++{
		for j :=i+1; j<10; j++{
			for k :=j+1; k<10; k++{
				if(i!=j && i!= k && j!=k){
					fmt.Print(strconv.Itoa(i)+strconv.Itoa(j)+strconv.Itoa(k)+", ")
				}
			}
		}
	}
}
