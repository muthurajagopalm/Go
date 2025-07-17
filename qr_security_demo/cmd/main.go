package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/muthurajagopalm/qr_security_demo/internal/sign"
	"github.com/muthurajagopalm/qr_security_demo/internal/verify"
)

func main() {
	fmt.Println("Choose an action:")
	fmt.Println("1. Generate QR Code")
	fmt.Println("2. Verify QR Code")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		err := sign.GenerateQRCode()
		if err != nil {
			fmt.Println("Error:", err)
		}
	case 2:
		fmt.Println("Paste QR JSON string:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		var qrData struct {
			Payload   string `json:"payload"`
			Signature string `json:"signature"`
		}

		err := json.Unmarshal([]byte(input), &qrData)
		if err != nil {
			fmt.Println("Invalid JSON:", err)
			return
		}

		payloadBytes, err := base64.StdEncoding.DecodeString(qrData.Payload)
		if err != nil {
			fmt.Println("Error decoding payload:", err)
			return
		}

		err = verify.VerifyQRCode(string(payloadBytes), qrData.Signature)
		if err != nil {
			fmt.Println("Verification failed:", err)
		}

	default:
		fmt.Println("Invalid choice.")
	}
}
