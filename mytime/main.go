package main
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("08:05 Monday"))

	createDate := time.Date(2020,time.January,24,10,20,20,20,time.UTC)
	fmt.Println(createDate)
}