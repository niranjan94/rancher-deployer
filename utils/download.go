package utils

import (
	"os"
	"net/http"
	"io"
	"github.com/mholt/archiver"
	"io/ioutil"
	"log"
)

// Download a file from a URL
func DownloadFile(filePath string, url string) error {
	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// Extract a .tar.gz archive
func Extract(filePath string) string  {
	destinationDirectory, _ := ioutil.TempDir(os.TempDir(), "go-deployer")
	err := archiver.TarGz.Open(filePath, destinationDirectory)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return  destinationDirectory
}

// Download and extract a file
func DownloadExtract(url string) string  {
	destinationFile, _ := ioutil.TempFile(os.TempDir(), "go-deployer")
	DownloadFile(destinationFile.Name(), url)
	destinationDirectory := Extract(destinationFile.Name())
	defer os.Remove(destinationFile.Name())
	return  destinationDirectory
}