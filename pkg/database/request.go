package database

import (
	"context"
	"github.com/Maden-in-haven/crmlib/pkg/model"
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

func GetUserByID(userID string) (*model.User, error) {
	// SQL запрос
	query := `
		SELECT id, username, password_hash, role, created_at, updated_at, is_deleted
		FROM users
		WHERE id = $1
	`

	// Создаем переменную для пользователя
	var user model.User

	// Выполняем запрос
	err := DbPool.QueryRow(context.Background(), query, userID).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.Role,
		&user.CreatedAt, &user.UpdatedAt, &user.IsDeleted,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func IsUserAdmin(userID string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM admins WHERE id = $1)`
	err := DbPool.QueryRow(context.Background(), query, userID).Scan(&exists)
	if err != nil {
		return false, err // в случае других ошибок
	}
	return exists, nil
}
