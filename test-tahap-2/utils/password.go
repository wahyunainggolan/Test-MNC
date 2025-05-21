package utils

import "golang.org/x/crypto/bcrypt"

func HashPIN(pin string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(pin), 14)
    return string(bytes), err
}

func CheckPIN(hashedPin, pin string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPin), []byte(pin))
}