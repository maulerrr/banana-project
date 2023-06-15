package test

import (
	"testing"
	"time"

	"github.com/maulerrr/banana/pkg/models"
)

func TestHashPassword(t *testing.T) {
	password := "password123"
	hashedPassword := HashPassword([]byte(password))

	if hashedPassword == "" {
		t.Error("HashPassword should return a non-empty string")
	}
}

func TestComparePasswords(t *testing.T) {
	password := "password123"
	hashedPassword := HashPassword([]byte(password))

	match := ComparePasswords(hashedPassword, password)
	if !match {
		t.Error("ComparePasswords should return true for matching passwords")
	}

	wrongPassword := "wrongpassword"
	match = ComparePasswords(hashedPassword, wrongPassword)
	if match {
		t.Error("ComparePasswords should return false for non-matching passwords")
	}
}

func TestGenerateTokenAndParseToken(t *testing.T) {
	jwtKey := "secretkey"
	user := models.User{
		UserID:    1,
		Username:  "testuser",
		Email:     "test@example.com",
		CreatedAt: time.Now(),
	}

	token, err := GenerateToken(user, jwtKey)
	if err != nil {
		t.Errorf("GenerateToken returned an error: %v", err)
	}

	claims, err := ParseToken(token, jwtKey)
	if err != nil {
		t.Errorf("ParseToken returned an error: %v", err)
	}

	if claims.ID != user.UserID {
		t.Errorf("Parsed token claims mismatch: expected ID %d, got %d", user.UserID, claims.ID)
	}

	if claims.Username != user.Username {
		t.Errorf("Parsed token claims mismatch: expected username %s, got %s", user.Username, claims.Username)
	}

	if claims.Email != user.Email {
		t.Errorf("Parsed token claims mismatch: expected email %s, got %s", user.Email, claims.Email)
	}
}
