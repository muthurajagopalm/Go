package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/liyue201/goqr"
)

func main() {
	//Step 1: Open the image file
	file, err := os.Open("qrcode.png")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	//Step 2: Decode the image
	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	fmt.Println("Image Sucessfully loaded")
	fmt.Println("Image format:", format)
	fmt.Println("Image type:", img.Bounds())

	//Step 3: Scan the image for QR code

	codes, err := goqr.Recognize(img)
	if err != nil {
		fmt.Println("Error recognizing QR code:", err)
		return
	}

	//step 4: Print the QR code data
	if len(codes) == 0 {
		fmt.Println("No QR code found in the image.")
	} else {
		for i, code := range codes {
			fmt.Printf("QR Code %d: %s\n", i+1, code.Payload)
		}
	}
}
