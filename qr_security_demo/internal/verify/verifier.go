package verify

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func VerifyQRCode(payloadJSON string, signatureBase64 string) error {
	signatureBytes, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return fmt.Errorf("error decoding signature: %v", err)
	}

	publicKey, err := loadPublicKeyFromFile("keys/public.pem")
	if err != nil {
		return fmt.Errorf("error loading public key: %v", err)
	}

	hashed := sha256.Sum256([]byte(payloadJSON))

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signatureBytes)
	if err != nil {
		fmt.Println("Signature verification FAILED!")
		return err
	}

	fmt.Println("Signature verification SUCCESSFUL ✅")
	return nil
}

func loadPublicKeyFromFile(filename string) (*rsa.PublicKey, error) {
	pubPEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pubPEM)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return publicKey, nil
}
