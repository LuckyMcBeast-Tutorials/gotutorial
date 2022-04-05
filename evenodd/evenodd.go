package main

import "fmt"

type oddEven []int

func newOddEven() oddEven{
	return oddEven{0,1,2,3,4,5,6,7,8,9,10}
}

func isEven(num int) bool {
	return num % 2 == 0
}

func oddOrEvenOutput(num int) string {
	if isEven(num){
		return fmt.Sprintf("%v is even", num)
	}
	return fmt.Sprintf("%v is odd", num)
}

func (o oddEven) print() {
	for _, num := range o {
		fmt.Println(oddOrEvenOutput(num))
	}
}