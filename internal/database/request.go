package database

import (
	"github.com/Maden-in-haven/crmlib/pkg/model"
	"context"
)

func GetUserByUsername(username string) (*model.User, error) {
	// SQL запрос
	query := `
		SELECT id, username, password_hash, role, created_at, updated_at, is_deleted
		FROM users
		WHERE username = $1
	`

	// Создаем переменную для пользователя
	var user model.User

	// Выполняем запрос
	err := DbPool.QueryRow(context.Background(), query, username).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.Role,
		&user.CreatedAt, &user.UpdatedAt, &user.IsDeleted,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
