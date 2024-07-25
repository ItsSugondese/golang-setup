package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	filepathconstants "wabustock/constants/file_path_constants"
	"wabustock/constants/file_type_constants"
	globaldto "wabustock/global/global_dto"
)

// SaveFile saves the uploaded file to the specified directory and returns the URL of the saved file.
func SaveFile(file *multipart.FileHeader, module string) globaldto.FileDetails {
	uploadDir := filepath.Join(filepathconstants.UploadDir, filepathconstants.FilePathMappings[module].Path)
	// Create the upload directory if it doesn't exist
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			panic("unable to create directory: " + err.Error())
		}
	}

	fileType := validateExtension(file.Filename)
	// Create a unique file name based on the current timestamp
	timestamp := time.Now().UnixNano()
	extension := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d%s", timestamp, extension)
	filePath := filepath.Join(uploadDir, newFileName)

	// SAVE the file to the specified directory
	if err := saveUploadedFile(file, filePath); err != nil {
		panic("unable to save the file: " + err.Error())
	}

	// Return the URL of the saved file
	//fileURL := "localhost:3000/images/" + newFileName
	return globaldto.FileDetails{
		FilePath: filePath,
		Size:     file.Size,
		FileType: fileType,
	}
}

// saveUploadedFile is a helper function to save the uploaded file to the file system
func saveUploadedFile(file *multipart.FileHeader, filePath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.ReadFrom(src)
	return err
}

// fins and return the file from the path
func GetFileFromFilePath(filePath string, w http.ResponseWriter) {
	if filePath == "" {
		panic("File path is required")

	}

	fileName := filepath.Base(filePath)
	if fileName == "" {
		panic("Invalid file name")

	}

	file, err := os.Open(filePath)
	if err != nil {
		panic("Invalid file path")

	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		panic("Failed to write file to response")

	}
}

// responsible for copying file from one path to another. will primiralrly be used to copy from temporary file to actual file path
func CopyFileToServer(filePath string, fileToPath string) string {
	fileName := filepath.Base(filePath)
	currentTime := time.Now()

	// Format the date as YYYY-MM-DD
	date := currentTime.Format("2006-01-02")
	fileTo := filepath.Join(filepathconstants.UploadDir, filepathconstants.FilePathMappings[fileToPath].Location, date)

	overallToPath := filepath.Join(fileTo, fileName)
	// Create the directory if it doesn't exist
	err := os.MkdirAll(fileTo, os.ModePerm)
	if err != nil {
		panic("Failed to create directory: " + err.Error())
	}

	// Copy the file
	err = copyFile(filePath, overallToPath)
	if err != nil {
		panic("Failed to copy the file: " + err.Error())
	}

	return overallToPath
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

// for validating whether the extension of the file is valid or not
func validateExtension(filename string) file_type_constants.FileType {
	extension := strings.ToUpper(filepath.Ext(filename))[1:] // get file extension without dot

	// Check if the extension is empty
	if extension == "" {
		panic("file has no extension")
	}

	var fileType file_type_constants.FileType

	if fileType, ok := file_type_constants.ImageType[extension]; ok {
		return fileType
	} else if fileType, ok := file_type_constants.DocumentType[extension]; ok {
		return fileType
	} else if fileType, ok := file_type_constants.PdfType[extension]; ok {
		return fileType
	} else if fileType, ok := file_type_constants.TxtType[extension]; ok {
		return fileType
	} else if fileType, ok := file_type_constants.ExcelType[extension]; ok {
		return fileType
	} else {
		panic("Not a valid extension")
	}

	// Prepare the result map
	return fileType
}
