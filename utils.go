package model

import "os"

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateFile(filename string) error {
	_, err := os.Create(filename)
	return err
}

func AppendToFile(filename, content string) error {
	// Open the file in append mode, create it if it doesn't exist
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write the content to the file
	if _, err := f.WriteString(content + "\n"); err != nil {
		return err
	}
	return nil
}
