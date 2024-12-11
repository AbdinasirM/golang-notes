package main
import "fmt"


func main(){
	nums:= []int{16,8,42,4,23,15}
	max :=nums[0]

	for _,num:=range nums{
		if num > max {
			max = num
		}
	}

	fmt.Println("The maximum value is ", max)
}