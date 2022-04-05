package main

import "fmt"

func main(){
	triangle := triangle{base: 4.20, height: 7.10}
	square := square{sideLen: 9.6}

	fmt.Println("The area of the triangle is", triangle.getArea())
	fmt.Println("The area of the square is", square.getArea())
}