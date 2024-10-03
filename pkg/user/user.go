package user

import (
	"context"

	"github.com/Maden-in-haven/crmlib/pkg/database"
	"github.com/Maden-in-haven/crmlib/pkg/model"
	"golang.org/x/crypto/bcrypt"
)

// Функция для проверки пароля
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Функция для аутентификации пользователя
func AuthenticateUser(username, password string) (model.User, error) {
	// Находим пользователя по username
	user, err := database.DB.GetUserByUsername(context.Background(), username)
	if err != nil {
		return model.User{}, err
	}

	// Проверяем пароль
	err = CheckPassword(user.PasswordHash, password)
	if err != nil {
		return model.User{}, err
	}

	// Возвращаем пользователя, если пароль верен
	return user, nil
}

