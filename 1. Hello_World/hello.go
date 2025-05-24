package main

import (
	"fmt"
	"math"
)

// func main(){
// 	fmt.Print("Hello "+"World!!",rand.Intn(10))
// }

// func main(){
// 	fmt.Print("Pi value is ", math.Pi)
// }

// func add(x int , y int) int {// if both are of type int you can write x, y int
// 	return x + y
// }

// func swap(x,y string) (string,string) {
// 	return y,x
// }
// func main(){
// 	// var a, b string or we can do this inside the function
// 	a, b := swap("Adi", "Cosmo")
// 	fmt.Println(a,b)
// }

// func split(sum int )(x,y int){
// 	x = sum*4/9
// 	y= sum-x

// 	return
// }

// func main(){
// 	fmt.Println(split(17))
// }

// var (
// 	Bool bool = false
// 	MaxInt uint64 = 1<<64-1
// 	z complex128 = cmplx.Sqrt(-5+12i)
// )

// func main(){
// 	fmt.Printf("Type : %T Value is %v\n", Bool,Bool)
// 	fmt.Printf("Type : %T Value is %v\n", MaxInt,MaxInt)
// 	fmt.Printf("Type : %T Value is %v\n", z,z)
// }

func main(){
	var x, y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Print(x,y,z)
}