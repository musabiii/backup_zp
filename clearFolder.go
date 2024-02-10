package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	// crear folder, remove files older then 7 days
	folderPath := "F:\\1c\\backup\\files"
	info, _ := ioutil.ReadDir(folderPath)
	for _, f := range info {
		if !f.IsDir() && (time.Since(f.ModTime()) > 7*24*time.Hour) { // 7 days
			os.Remove(fmt.Sprintf("%s\\%s", folderPath, f.Name()))
		}
	}
}
