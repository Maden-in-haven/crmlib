package user

import (
	"context"

	"github.com/Maden-in-haven/crmlib/pkg/database"
	"github.com/Maden-in-haven/crmlib/pkg/model"
	"github.com/Maden-in-haven/crmlib/pkg/util"
)

// Функция для аутентификации пользователя
func AuthenticateUser(username, password string) (model.User, error) {
	// Находим пользователя по username
	user, err := database.DB.GetUserByUsername(context.Background(), username)
	if err != nil {
		return model.User{}, err
	}

	// Проверяем пароль
	err = util.CheckPassword(user.PasswordHash, password)
	if err != nil {
		return model.User{}, err
	}

	// Возвращаем пользователя, если пароль верен
	return user, nil
}
