package imageUploader

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/oldgattsu/diplom2/internal/config"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

func Load(rawImage string) (string, error) {
	index := strings.Index(rawImage, "base64,")
	imageBase64, _ := b64.StdEncoding.DecodeString(rawImage[index+7:])
	reader := bytes.NewReader(imageBase64)

	image, format, imageDecodeErr := image.Decode(reader)
	if imageDecodeErr != nil {
		fmt.Printf("error: %v", imageDecodeErr)
		return "", imageDecodeErr
	}
	fmt.Printf("image format: %v", format)

	dir := "static"
	uniqueFileName := uuid.New().String()
	fullFileName := fmt.Sprintf("%s/%s.%s", dir, uniqueFileName, format)

	f, openFilerErr := os.OpenFile(fullFileName, os.O_WRONLY|os.O_CREATE, 0777)
	if openFilerErr != nil {
		fmt.Printf("error: %v", openFilerErr)
		return "", openFilerErr
	}

	switch format {
	case "png":
		png.Encode(f, image)
	case "jpeg":
		jpeg.Encode(f, image, nil)
	}

	cfg, errCfg := config.Load()
	if errCfg != nil {
		log.Printf("error load config, %v", errCfg)
	}

	imageURL := fmt.Sprintf("http://%s/%s", cfg.Address, fullFileName)

	return imageURL, nil
}
