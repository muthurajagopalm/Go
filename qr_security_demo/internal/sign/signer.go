package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode() error {
	// Create payload
	payload := map[string]interface{}{
		"merchant_id": "M12345",
		"amount":      250,
		"currency":    "INR",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	fmt.Println("Payload JSON:")
	fmt.Println(string(payloadBytes))
	fmt.Println("")

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// Save private key
	err = savePrivateKeyToFile(privateKey, "keys/private.pem")
	if err != nil {
		return err
	}

	// Save public key
	publicKey := &privateKey.PublicKey
	err = savePublicKeyToFile(publicKey, "keys/public.pem")
	if err != nil {
		return err
	}

	hashed := sha256.Sum256(payloadBytes)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return err
	}

	payloadBase64 := base64.StdEncoding.EncodeToString(payloadBytes)
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	secureData := map[string]string{
		"payload":   payloadBase64,
		"signature": signatureBase64,
	}

	secureDataJSON, err := json.Marshal(secureData)
	if err != nil {
		return err
	}

	err = qrcode.WriteFile(string(secureDataJSON), qrcode.Medium, 256, "qrcode.png")
	if err != nil {
		return err
	}

	fmt.Println("QR code generated as qrcode.png ✅")

	return nil
}

func savePrivateKeyToFile(privateKey *rsa.PrivateKey, filename string) error {
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)
	privBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privDER,
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, privBlock)
}

func savePublicKeyToFile(publicKey *rsa.PublicKey, filename string) error {
	pubDER, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubDER,
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, pubBlock)
}
