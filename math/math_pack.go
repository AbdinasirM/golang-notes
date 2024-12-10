package main
import (
	"fmt"
	"math"
	

)
func main(){
	

	i1, i2, i3  := 11,43,68;
	int_sum := i1 +i2 +i3
	fmt.Println("the total is:", int_sum)
	total := math.Round(float64(int_sum))
	fmt.Println("the total is:", total)
}