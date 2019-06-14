package piscine_go

import "math"

import "fmt"

func IterativeFactorial(nb int) int{
	s:=0
	for i:=1; i<=nb; i++{
		s*= i
	}
	return s
}

func RecursiveFactorial(nb int) int{
	if nb <=1{
		return 1
	} else{
		return nb * RecursiveFactorial(nb -1)
	}
}

func IterativePower(nb, power int) int{
	p:= 1
	for i:=1; i<=power; i++{
		p*= nb
	}
	return p
}

func RecursivePower(nb, power int) int{
	if power<=0{
		return 0
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}

func Fibonacci(index int) int {
	if index<2 {
		return index
	} else {
		return Fibonacci(index -1) + Fibonacci(index -2)
	}
}

func Sqrt(nb int) int {
	aux:=math.Sqrt(float64(nb))
	aux2 :=int(aux)
	if float64(aux2) == aux{
		return int(aux)
	} else{
		return 0
	}
}

func IsPrime(nb int) bool {
	for i := 2; i <= int(math.Floor(float64(nb) / 2)); i++ {
		if nb%i == 0 {
			return false
		}
	}
	return nb > 1
}

func FindNextPrime(nb int) int {
	for !IsPrime(nb){
		nb+=1
	}
	return nb
}

type Point struct {
	x int
	y int
}

var results = make([][]Point, 0)


func EightQueens() {
	Solve(8)

}

func Solve(n int) {
	for col := 0; col < n; col++ {
		start := Point{x: col, y: 0}
		current := make([]Point, 0)
		Recurse(start, current, n)
	}
	for _, result := range results {
		for _,point := range result{
			fmt.Print(point.x+1)
		}
		fmt.Print("\n")
	}
}
func Recurse(point Point, current []Point, n int) {
	if CanPlace(point, current) {
		current = append(current, point)
		if len(current) == n {
			c := make([]Point, n)
			for i, point := range current {
				c[i] = point
			}
			results = append(results, c)
		} else {
			for col := 0; col < n; col++ {
				for row := point.y; row < n; row++ {
					nextStart := Point{x: col, y: row}
					Recurse(nextStart, current, n)

				}

			}
		}
	}
}
func CanPlace(target Point, board []Point) bool {
	for _, point := range board {
		if CanAttack(point, target) {
			return false
		}
	}
	return true
}

func CanAttack(a, b Point) bool {
	//fmt.Print(a, b)
	answer := a.x == b.x || a.y == b.y || math.Abs(float64(a.y-b.y)) == math.Abs(float64(a.x-b.x))
	//fmt.Print(answer)
	return answer
}