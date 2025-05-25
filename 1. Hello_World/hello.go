package main

import "fmt"

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

// func main(){
// 	var x, y int = 3,4
// 	var f float64 = math.Sqrt(float64(x*x + y*y))
// 	var z uint = uint(f)

// 	fmt.Print(x,y,z)
// }

// func main(){
// 	var username string= "Aditya"
// 	var password string = "12345"
// 	fmt.Println("Authorization: Basic", username+":"+password)
// }

// func main(){
// 	var smsSendingLimit int
// 	var costPerSms float64
// 	var hasPermission bool
// 	var username string

// 	fmt.Printf("%v %f %v %q\n", smsSendingLimit,costPerSms,hasPermission,username,)
// }

// func main(){
// 	// var empty string
// 	// empty2 := ""
// 	// they both are sae thing

// 	congrats := "Happy birthday!!"
// 	fmt.Println(congrats)
// }

// func main(){
// 	penniesPerText := 2.0
// 	fmt.Printf("The type of the variable is : %T\n",penniesPerText)
// }

// func main(){
// 	avgOpenRate , displayMessage := 0.23, "message"

// 	fmt.Printf("Types are %T,%T and values are %v,%v\n",avgOpenRate,displayMessage,avgOpenRate,displayMessage )

// }

// const did not allow the short declaration syntax

// func main(){
// 	const a = 5
// 	const b = a+3

// 	var x = 10
// 	const y = x +1// error : as const should be known in complite time

// 	fmt.Print(y)
// }

func main(){
	name := "Aditya"
	openRate := 20.99

	msg := fmt.Sprintf("Hi %s your open rate is %.1f percent", name,openRate)

	fmt.Println(msg)
}