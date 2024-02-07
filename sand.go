package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")
	currentTime := time.Now()
	fmt.Println(currentTime.Format("02_01_2006"))
}
