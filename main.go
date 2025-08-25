package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

type PasswordData struct {
	PlainText string `json:"plain_text"`
	Salt      string `json:"salt"`
	Hashed    string `json:"hashed"`
}

// generateRandomSalt 길이 length만큼의 랜덤 salt를 생성하는 함수
func generateRandomSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

// hashPlainTextWithSaltAndStretching 평문에 salt를 붙이고, 해시를 반복 계산하는 함수
func hashPlainTextWithSaltAndStretching(plainText, salt string, iterations int) string {
	data := []byte(plainText + salt)
	var hash []byte

	for i := 0; i < iterations; i++ {
		h := sha256.New()
		h.Write(data)
		hash = h.Sum(nil)
		data = hash // 다음 반복 때 해시값을 입력으로 사용
	}

	return hex.EncodeToString(hash)
}

func saveToFile(filename string, data PasswordData) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the original data: ")
	if !scanner.Scan() {
		fmt.Println("No input detected.")
		return
	}
	plainText := scanner.Text()
	saltLength := 16 // 16바이트 salt 생성 (32자리 16진수 문자열)
	iterations := 10000

	salt, err := generateRandomSalt(saltLength)
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	hashed := hashPlainTextWithSaltAndStretching(plainText, salt, iterations)
	fmt.Println("Salt:", salt)
	fmt.Println("Hashed value:", hashed)

	/*if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}*/
	data := PasswordData{
		PlainText: plainText,
		Salt:      salt,
		Hashed:    hashed,
	}

	filename := "password_data.json"
	if err := saveToFile(filename, data); err != nil {
		fmt.Printf("Error saving to file: %v\n", err)
		return
	}
	fmt.Println("Success to save data")
}
