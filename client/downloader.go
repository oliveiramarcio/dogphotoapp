package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func downloadAndSavePhoto(photoDir, photoURL string) error {
	resp, err := http.Get(photoURL)
	if err != nil {
		return fmt.Errorf("failed to download photo: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download photo: %s", resp.Status)
	}

	err = os.MkdirAll(photoDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create photo directory: %v", err)
	}

	urlChecksum := calculateMD5(photoURL)

	filename := fmt.Sprintf("%s.jpg", urlChecksum)
	savePath := filepath.Join(photoDir, filename)

	if _, err := os.Stat(savePath); err == nil {
		log.Printf("file already exists: %s", savePath)
		return nil
	}

	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("failed to create photo file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save photo: %v", err)
	}

	log.Printf("dog photo saved: %s", photoURL)

	return nil
}
