package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	cleartext := []byte("The quick brown fox jumps over the lazy dog.")

	password := []byte("mysuperdupersecurepassword")
	salt := NewNonce()
	key := NewKey(salt, password)
	nonce := NewNonce()
	ciphertext, err := Encrypt(key, nonce, cleartext)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(ciphertext, cleartext) {
		t.Fatal("Ciphertext and cleartext are the same.")
	}

	decrypted, err := Decrypt(key, nonce, ciphertext)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(cleartext, decrypted) {
		t.Fatal("Original cleartext and decrypted version are different.")
	}
}

func TestEncrypt(t *testing.T) {
	// 关键参数应该是AES密钥，16或32个字节
	// 选择AES-128或AES-256。
	key := []byte("AES256Key-32Characters1234567890")
	plaintext := []byte("exampleplaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// 由于存在重复的风险，请勿使用给定密钥使用超过2^32个随机值。
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("%x\n", ciphertext)
}
func TestDecrypt(t *testing.T) {
	// 关键参数应该是AES密钥，16或32个字节
	// 选择 AES-128 或 AES-256。
	key := []byte("AES256Key-32Characters1234567890")
	ciphertext, _ := hex.DecodeString("1019aa66cd7c024f9efd0038899dae1973ee69427f5a6579eba292ffe1b5a260")

	nonce, _ := hex.DecodeString("37b8e8a308c354048d245f6d")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", plaintext)
}

const BUFFER_SIZE int = 4096
const IV_SIZE int = 16

func encrypt(filePathIn, filePathOut string, keyAes, keyHmac []byte) error {
	inFile, err := os.Open(filePathIn)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(filePathOut)
	if err != nil {
		return err
	}
	defer outFile.Close()

	iv := make([]byte, IV_SIZE)
	_, err = rand.Read(iv)
	if err != nil {
		return err
	}

	aes, err := aes.NewCipher(keyAes)
	if err != nil {
		return err
	}

	ctr := cipher.NewCTR(aes, iv)
	hmac := hmac.New(sha256.New, keyHmac)

	buf := make([]byte, BUFFER_SIZE)
	for {
		n, err := inFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		outBuf := make([]byte, n)
		ctr.XORKeyStream(outBuf, buf[:n])
		hmac.Write(outBuf)
		outFile.Write(outBuf)

		if err == io.EOF {
			break
		}
	}

	outFile.Write(iv)
	hmac.Write(iv)
	outFile.Write(hmac.Sum(nil))

	return nil
}
