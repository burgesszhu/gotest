package add

import "fmt"

func Add(a,b float64) float64{
	c := a + b
	fmt.Sprintln(a, b, c)
	return c

}

