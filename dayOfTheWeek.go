package main

import (
	"fmt"
	"time"
)

func main() {
	//Insert code here
	switch time.Now().Weekday() {
	case time.Monday, time.Wednesday, time.Friday:
		fmt.Println("Today is an odd day")
	case time.Tuesday, time.Thursday:
		fmt.Println("Today is an odd day")
	default:
		fmt.Println("It is a week day")
	}
}
