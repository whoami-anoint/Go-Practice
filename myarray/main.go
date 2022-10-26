package main
import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

var fruitlist[4]string
fruitlist[0] = "Apple"
fruitlist[1] = "Orange"
fruitlist[2] = "Banana"
fruitlist[3] = "Mango"
fmt.Println(fruitlist)
fmt.Println(len(fruitlist))

var veglist = [3]string{"Potato","Onion","Beans"}
fmt.Println(veglist)
fmt.Println(len(veglist))

}