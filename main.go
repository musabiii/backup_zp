package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"time"
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

	currentTime := time.Now()
	dateString := currentTime.Format("02_01_2006")

	archivePath := fmt.Sprintf("F:\\1c\\backup\\files\\%s.zip", dateString)
	dbFile := "\\\\servhp\\Base\\zup2\\1Cv8.1CD"

	// 2. Archive the file to a zip file
	zipFile, err := os.Create(archivePath)
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fileToArchive, err := os.Open(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fileToArchive.Close()

	fileInZip, err := zipWriter.Create("1Cv8.1CD")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(fileInZip, fileToArchive)
	if err != nil {
		log.Fatal(err)
	}
}
