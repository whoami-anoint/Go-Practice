package main

import (
	"fmt"
	"math/big"
	//"math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	// 	fmt.Println("Welcome to math in golang")

	// 	// var mynumberOne int = 10
	// 	// var mynumberTwo float64 = 4.5
	// 	// fmt.Println("The sum of two numbers is",mynumberOne + int(mynumberTwo))
	// }

	// random number
// 	rand.Seed(time.Now().UnixNano())
// 	fmt.Println(rand.Intn(5)+1)

// }
 myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(24))  
fmt.Println(myRandomNum)
}
