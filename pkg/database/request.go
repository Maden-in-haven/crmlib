package database

import (
	"context"
	"fmt"
	"time"

	"github.com/Maden-in-haven/crmlib/pkg/model"
	"github.com/jackc/pgx/v5"
)

// GetAllUsers возвращает список всех пользователей из таблицы users, у которых флаг is_deleted = false.
func (db *db) GetAllUsers(ctx context.Context) ([]model.User, error) {
	query := `SELECT id, username, role, created_at, updated_at FROM users WHERE is_deleted = false`

	rows, err := db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}

	for rows.Next() {
		var user model.User
		var createdAt time.Time
		var updatedAt time.Time

		err := rows.Scan(&user.ID, &user.Username, &user.Role, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		user.CreatedAt = createdAt.Format(time.RFC3339)
		user.UpdatedAt = updatedAt.Format(time.RFC3339)
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (db *db) GetUserByID(ctx context.Context, userID string) (model.User, error) {
	query := `SELECT id, username, role, password_hash, created_at, updated_at 
			  FROM users 
			  WHERE id = $1 AND is_deleted = false`

	var user model.User
	var createdAt time.Time
	var updatedAt time.Time

	// Выполнение SQL-запроса для получения пользователя по имени пользователя
	err := db.Pool.QueryRow(ctx, query, userID).Scan(&user.ID, &user.Username, &user.Role, &user.PasswordHash, &createdAt, &updatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user, fmt.Errorf("пользователь с ID %s не найден", userID)
		}
		return user, err
	}

	// Преобразуем временные метки в строку
	user.CreatedAt = createdAt.Format(time.RFC3339)
	user.UpdatedAt = updatedAt.Format(time.RFC3339)

	return user, nil
}

func (db *db) GetAdminByID(ctx context.Context, adminID string) (model.Admin, error) {
	query := `SELECT u.id, u.username, a.permissions, u.created_at 
			  FROM admins a 
			  JOIN users u ON a.id = u.id 
			  WHERE u.id = $1 AND u.is_deleted = false`

	var admin model.Admin
	var createdAt time.Time

	// Выполнение SQL-запроса для получения администратора по ID
	err := db.Pool.QueryRow(ctx, query, adminID).Scan(&admin.ID, &admin.Username, &admin.Permissions, &createdAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return admin, fmt.Errorf("администратор с ID %s не найден", adminID)
		}
		return admin, err
	}

	// Преобразуем временные метки в строку
	admin.CreatedAt = createdAt.Format(time.RFC3339)

	return admin, nil
}

func (db *db) GetClientByID(ctx context.Context, clientID string) (model.Client, error) {
	query := `SELECT u.id, u.username, c.full_name, c.phone_number, u.created_at 
			  FROM clients c 
			  JOIN users u ON c.id = u.id 
			  WHERE u.id = $1 AND u.is_deleted = false`

	var client model.Client
	var createdAt time.Time

	// Выполнение SQL-запроса для получения клиента по ID
	err := db.Pool.QueryRow(ctx, query, clientID).Scan(&client.ID, &client.Username, &client.FullName, &client.PhoneNumber, &createdAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return client, fmt.Errorf("клиент с ID %s не найден", clientID)
		}
		return client, err
	}

	// Преобразуем временные метки в строку
	client.CreatedAt = createdAt.Format(time.RFC3339)

	return client, nil
}

func (db *db) GetManagerByID(ctx context.Context, managerID string) (model.Manager, error) {
	query := `SELECT u.id, u.username, m.full_name, m.hire_date, u.created_at 
			  FROM managers m 
			  JOIN users u ON m.id = u.id 
			  WHERE u.id = $1 AND u.is_deleted = false`

	var manager model.Manager
	var createdAt time.Time
	var hireDate time.Time

	// Выполнение SQL-запроса для получения менеджера по ID
	err := db.Pool.QueryRow(ctx, query, managerID).Scan(&manager.ID, &manager.Username, &manager.FullName, &hireDate, &createdAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return manager, fmt.Errorf("менеджер с ID %s не найден", managerID)
		}
		return manager, err
	}

	// Преобразуем временные метки в строку
	manager.CreatedAt = createdAt.Format(time.RFC3339)
	manager.HireDate = hireDate.Format(time.RFC3339)

	return manager, nil
}

func (db *db) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	query := `SELECT id, username, role, password_hash, created_at, updated_at 
			  FROM users 
			  WHERE username = $1 AND is_deleted = false`

	var user model.User
	var createdAt time.Time
	var updatedAt time.Time

	// Выполнение SQL-запроса для получения пользователя по имени пользователя
	err := db.Pool.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.Role, &user.PasswordHash, &createdAt, &updatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user, fmt.Errorf("пользователь с именем %s не найден", username)
		}
		return user, err
	}

	// Преобразуем временные метки в строку
	user.CreatedAt = createdAt.Format(time.RFC3339)
	user.UpdatedAt = updatedAt.Format(time.RFC3339)

	return user, nil
}