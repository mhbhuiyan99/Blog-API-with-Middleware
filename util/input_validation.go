package util

import (
	"fmt"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func ValidateEmail(email string) error {

	email = strings.TrimSpace(email)
	if email == "" {
		return fmt.Errorf("Email cannot be empty")
	}

	if !emailRegex.MatchString(email) {
		return fmt.Errorf("Invalid email format")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("Password must be at least 8 characters long")
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	if !hasUpper || !hasLower || !hasNumber {
		return fmt.Errorf("Password must contain at least one uppercase letter, one lowercase letter, and one number")
	}

	return nil
}

func ValidateUsername(username string) error {
	username = strings.TrimSpace(username)

	if len(username) < 3 || len(username) > 20 {
		return fmt.Errorf("username must be between 3 and 20 characters long")
	}

	validUsername := regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(username)
	if !validUsername {
		return fmt.Errorf("username can only contain letters, numbers, and underscores")
	}

	return nil
}

func ValidateTitle(title string) error {
	title = strings.TrimSpace(title)

	if title == "" {
		return fmt.Errorf("Title is required")
	}

	if len(title) < 3 || len(title) > 100 {
		return fmt.Errorf("Title must be between 3 and 100 characters long")
	}

	return nil
}

