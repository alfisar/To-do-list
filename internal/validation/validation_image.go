package validations

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"todolist/config"
)

func CheckDirectory(dirname string) (err error) {
	_, err = os.Stat(dirname)
	if err == nil {
		return
	}

	if os.IsNotExist(err) {
		err = os.MkdirAll(dirname, 0750) // 0750 is the permission for the directory
		return
	}

	return
}

func SaveImage(config *config.Config, DirImage string, fileHeader *multipart.FileHeader) (name string, err error) {

	var (
		pattern  string
		tempFile *os.File
		// fileBytes  []byte
		file multipart.File

		slicedName []string
	)

	err = CheckDirectory(DirImage)
	if err != nil {
		return
	}

	file, _ = fileHeader.Open()
	defer file.Close()

	switch fileHeader.Header.Values("Content-Type")[0] {

	case imageJPG, imageJPEG, imagePNG: // continue
	default:
		return name, errInvalidImageType
	}

	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	ext := filepath.Ext(fileHeader.Filename)
	validExtension := false

	for _, allowedExt := range allowedExtensions {
		if strings.EqualFold(ext, allowedExt) {
			validExtension = true
			break
		}
	}

	if !validExtension {
		return name, errInvalidImageType
	}

	isImage := isImageFile(fileHeader)

	if !isImage {
		return name, errInvalidImageType
	}

	removeSpace := strings.ReplaceAll(fileHeader.Filename, " ", "")
	pattern = fmt.Sprintf("*_%s", removeSpace)
	tempFile, err = os.CreateTemp(DirImage, pattern)

	if err != nil {
		return name, err
	}
	tempFile.Chmod(0755)
	defer tempFile.Close()

	slicedName = strings.Split(tempFile.Name(), "/")

	io.Copy(tempFile, file)

	// return tempFile.Name(), nil
	return slicedName[len(slicedName)-1], nil
}

func isImageFile(fileHeader *multipart.FileHeader) bool {
	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false
	}

	fileType := http.DetectContentType(buffer)
	return strings.HasPrefix(fileType, "image/")
}
