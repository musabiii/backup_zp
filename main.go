package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {

	currentTime := time.Now()
	dateString := currentTime.Format("02_01_2006")

	archivePath := fmt.Sprintf("F:\\1c\\backup\\files\\%s.zip", dateString)
	dbFile := "\\\\servhp\\Base\\zup2\\1Cv8.1CD"
	archiveDb(archivePath, dbFile)
	clearFolder()
}

func archiveDb(archivePath string, dbFile string) {

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

func clearFolder() {
	// crear folder, remove files older then 7 days
	folderPath := "F:\\1c\\backup\\files"
	info, _ := ioutil.ReadDir(folderPath)
	for _, f := range info {
		if !f.IsDir() && (time.Since(f.ModTime()) > 7*24*time.Hour) { // 7 days
			os.Remove(fmt.Sprintf("%s\\%s", folderPath, f.Name()))
		}
	}
}
