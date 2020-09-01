package util

import (
	"biller-api/src/api/constant"
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"os"

	"github.com/ventu-io/go-shortid"
)

func GenerateImage(b string) (string, error) {
	unbased, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		panic("Cannot decode b64")
	}

	r := bytes.NewReader(unbased)
	im, err := png.Decode(r)
	if err != nil {
		return "", fmt.Errorf("Error base64 reader: " + err.Error())
	}

	shortid, _ := shortid.Generate()
	path := constant.ImagePath + shortid + ".png"

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return "", fmt.Errorf("Error open file: " + err.Error())
	}

	err = png.Encode(f, im)
	if err != nil {
		return "", fmt.Errorf("Error create file: " + err.Error())
	}
	return path, nil
}
