package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	// "time"
)

func main() {
	// // 1. Copy a file to a folder
	// sourceFile, err := os.Open("./db/base.1CD")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("./backup/base.1CD")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer destFile.Close()

	// _, err = io.Copy(destFile, sourceFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// currentTime := time.Now()
	// fileName := currentTime.Format("02_01_2006")

	// formattedString := fmt.Sprintf("Hello, %s!", "World")

	// dbFile := "\\servhp\\Base\\zup2"

	// 2. Archive the file to a zip file
	zipFile, err := os.Create("./backup\\base.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fileToArchive, err := os.Open("./db/base.1CD")
	if err != nil {
		log.Fatal(err)
	}
	defer fileToArchive.Close()

	fileInZip, err := zipWriter.Create("base.1CD")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(fileInZip, fileToArchive)
	if err != nil {
		log.Fatal(err)
	}
}
