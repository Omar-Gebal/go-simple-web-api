package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Print("Enter password to hash: ")
	reader := bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading password: %v", err)
	}
	password = strings.TrimSpace(password)
	hashedPassword, err := hashPassword(password)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	// Insert backslashes before dollar signs
	formattedHash := strings.ReplaceAll(hashedPassword, "$", "\\$")

	fmt.Printf("Hashed password not formatted =%s\n", hashedPassword)
	fmt.Printf("Hashed password formatted for .env(to be stored in HASHED_AUTH_PASSWORD) =%s\n", formattedHash)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
