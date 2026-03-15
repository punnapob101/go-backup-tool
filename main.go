package main

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// List of files or folders to skip
var ignoreList = []string{".git", "node_modules", ".exe", ".zip", "__pycache__", ".DS_Store"}

func main() {
	var source string
	var destination string

	fmt.Println("==========================================")
	fmt.Println("🛡️  GO BACKUP TOOL")
	fmt.Println("==========================================")

	// 1. Get Source Path
	fmt.Print("📂 Enter folder to backup: ")
	fmt.Scanln(&source)
	source = strings.Trim(source, "\"")

	// Validate Source
	sourceInfo, err := os.Stat(source)
	if os.IsNotExist(err) || !sourceInfo.IsDir() {
		fmt.Println("❌ Error: Source folder does not exist.")
		return
	}

	// 2. Get Destination Path
	fmt.Print("📍 Where to save the backup? (e.g., D:/Backups): ")
	fmt.Scanln(&destination)
	destination = strings.Trim(destination, "\"")

	// Validate Destination (Create if it doesn't exist)
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		fmt.Printf("ℹ️ Destination folder not found. Creating: %s\n", destination)
		os.MkdirAll(destination, 0755)
	}

	// 3. Setup Target Filename
	folderName := filepath.Base(source)
	timestamp := time.Now().Format("2006-01-02_150405")
	targetFileName := fmt.Sprintf("%s_%s.zip", folderName, timestamp)
	targetPath := filepath.Join(destination, targetFileName)

	fmt.Printf("\n📦 Starting backup for: [%s]\n", folderName)
	fmt.Printf("🚀 Saving to: %s\n", targetPath)
	start := time.Now()

	// 4. Start Compression
	err = zipFolder(source, targetPath)
	if err != nil {
		fmt.Printf("❌ ERROR: %v\n", err)
		return
	}

	// 5. Check Integrity
	hash, _ := calculateSHA256(targetPath)

	fmt.Println("\n------------------------------------------")
	fmt.Println("✅ Backup Completed Successfully!")
	fmt.Printf("⏱️  Time Elapsed: %v\n", time.Since(start))
	fmt.Printf("📁 Filename: %s\n", targetFileName)
	fmt.Printf("🔒 SHA-256: %s\n", hash)
	fmt.Println("------------------------------------------")

	fmt.Println("------------------------------------------")
	fmt.Println("Created by Punnapob101 | github.com/Punnapob101")
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}

func zipFolder(source, target string) error {
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, ignore := range ignoreList {
			if strings.Contains(path, ignore) {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name, _ = filepath.Rel(filepath.Dir(source), path)
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})
}

func calculateSHA256(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
