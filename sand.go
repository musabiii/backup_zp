package main

import (
	"fmt"
	"time"
)

func main() {
	createFile()
}

func createFile() {
	//create file with specify modified date

}

func formatFunc() {
	fmt.Println("hello world")
	currentTime := time.Now()
	dateString := currentTime.Format("02_01_2006")
	archivePath := fmt.Sprintf("\\\\F:\\1c\\backup\\files\\%s.zip", dateString)
	fmt.Println(archivePath)
}
