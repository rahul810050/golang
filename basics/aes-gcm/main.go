package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)


func encrypt(data, key[]byte)([]byte, error){
	// this checks the length of the key is correct or not 
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// this will wrap aes with GCM(Glaois/Counter Mode) it will encrypt authenticate AES
	gcm, err := cipher.NewGCM(block) 
	if err != nil{
		return nil, err
	}
	// create nonce = "number used once"
	// Initial nonce (before random fill): nonce = [0 0 0 0 0 0 0 0 0 0 0 0]
	nonce := make([]byte, gcm.NonceSize()) // gcm.NonceSize() == 12
	_, err = rand.Read(nonce) // fill nonce with secure randomness nonce = [34 91 203 17 88 4 199 231 19 201 76 9]
	if err != nil{
		return nil, err
	}
	// encrypt + authenticate
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	// dst        = nonce
	// nonce      = [34 91 203 ...]
	// plaintext  = "PAYMENT=1000&USER=rahul"
	// AAD        = nil
	// internally:
	// GCM derives counters using nonce
	// AES encrypts plaintext
	// Authentication tag is generated
	// Everything is appended to dst


	return cipherText, nil

	// final encrypted output layout
	// ciphertext =[ NONCE (12 bytes) | ENCRYPTED DATA (len(data)) | AUTH TAG (16 bytes) ]
}

func decrypt(cipherText, key[]byte)([]byte, error){
	// recreate AES cipher with the same key same algorithm
	block, err := aes.NewCipher(key)
	if err != nil{
		return nil, err
	}
	// recreate GCM 
	gcm, err := cipher.NewGCM(block) // now decryptor is ready
	if err != nil {
		return nil, err
	}
	// get nonce size
	nonceSize := gcm.NonceSize()
	// extract nonce and encrypted payload
	// nonce = first 12 bytes
	// encrypted = encrypted data + auth tag
	nonce, encrypted := cipherText[:nonceSize], cipherText[nonceSize:]
	// authentication happens 
	return gcm.Open(nil, nonce, encrypted, nil)
}

func main() {
	// 32-byte key (AES-256)
	// In production: NEVER hardcode like this
	// Comes from ENV, Vault, KMS, Secret Manager
	key := []byte("12345678901234567890123456789012")

	// Plaintext data
	// This could be: password, token, PII, payment payload
	data := []byte("PAYMENT=1000&USER=rahul")

	fmt.Println("Original data:")
	fmt.Println(string(data))
	fmt.Println()

	// Encrypt
	encryptedData, err := encrypt(data, key)
	if err != nil {
		panic(err)
	}

	fmt.Println("Encrypted data (raw bytes):")
	fmt.Println(encryptedData)
	fmt.Println()

	// Encrypted data is usually stored/transmitted as base64
	encoded := base64.StdEncoding.EncodeToString(encryptedData)
	fmt.Println("Encrypted data (base64):")
	fmt.Println(encoded)
	fmt.Println()

	// Decrypt
	decryptedData, err := decrypt(encryptedData, key)
	if err != nil {
		panic(err)
	}

	fmt.Println("Decrypted data:")
	fmt.Println(string(decryptedData))
}
