package printprogramname

import (
	"fmt"
	"os"
)

func main()  {
	aux := os.Args
	fmt.Printf(aux[0])
}