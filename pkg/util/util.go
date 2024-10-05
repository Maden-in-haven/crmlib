package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Используем bcrypt для хеширования пароля
	// bcrypt.DefaultCost задает стандартную сложность хеширования
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Возвращаем хеш в строковом виде
	return string(hashedPassword), nil
}

// Функция для проверки пароля
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}