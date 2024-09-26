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

func IsUserAdmin(userID string) (exists bool) {
	query := `SELECT EXISTS (SELECT 1 FROM admins WHERE id = $1)`
	DbPool.QueryRow(context.Background(), query, userID).Scan(&exists)
	return exists
}

func GetUsersByRole(role string) ([]model.User, error) {
	var users []model.User

	// SQL-запрос для поиска пользователей по роли
	query := `
		SELECT id, username, role, created_at, updated_at, is_deleted
		FROM users
		WHERE role = $1
	`

	// Выполнение запроса
	rows, err := DbPool.Query(context.Background(), query, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Проход по строкам результата и сканирование данных в структуру
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.IsDeleted); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Проверка на ошибки, возникшие во время прохода по строкам
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
