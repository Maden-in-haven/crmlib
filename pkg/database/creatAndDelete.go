package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Maden-in-haven/crmlib/pkg/util"
	"time"
)

func (db *db) CreateAdmin(ctx context.Context, username, password string, permissions map[string]interface{}) (string, error) {
	query := `SELECT create_admin($1, $2, $3)`

	var adminID string

	// Преобразование карты permissions в JSONB формат
	permissionsJSON, err := json.Marshal(permissions)
	if err != nil {
		return "", fmt.Errorf("ошибка преобразования permissions в JSON: %v", err)
	}
	passwordHash, _ := util.HashPassword(password)
	// Выполнение запроса для вызова хранимой функции
	err = db.Pool.QueryRow(ctx, query, username, passwordHash, permissionsJSON).Scan(&adminID)
	if err != nil {
		return "", fmt.Errorf("ошибка вызова хранимой функции create_admin: %v", err)
	}

	// Логирование действия
	err = db.LogAction(ctx, adminID, fmt.Sprintf("Администратор %s был создан", username))
	if err != nil {
		return "", fmt.Errorf("ошибка записи лога: %v", err)
	}

	// Возвращаем ID нового администратора
	return adminID, nil
}

func (db *db) CreateClient(ctx context.Context, username, password, fullName, phoneNumber string) (string, error) {
	// SQL-запрос для вызова хранимой функции create_client
	query := `SELECT create_client($1, $2, $3, $4)`

	var clientID string
	passwordHash, _ := util.HashPassword(password)
	// Выполнение запроса для вызова хранимой функции
	err := db.Pool.QueryRow(ctx, query, username, passwordHash, fullName, phoneNumber).Scan(&clientID)
	if err != nil {
		return "", fmt.Errorf("ошибка вызова хранимой функции create_client: %v", err)
	}

	// Логирование действия
	err = db.LogAction(ctx, clientID, fmt.Sprintf("Клиент %s был создан", username))
	if err != nil {
		return "", fmt.Errorf("ошибка записи лога: %v", err)
	}

	// Возвращаем ID нового клиента
	return clientID, nil
}

func (db *db) CreateManager(ctx context.Context, username, password, fullName string, hireDate time.Time) (string, error) {
	// SQL-запрос для вызова хранимой функции create_manager
	query := `SELECT create_manager($1, $2, $3, $4)`

	var managerID string
	passwordHash, _ := util.HashPassword(password)
	// Выполнение запроса для вызова хранимой функции
	err := db.Pool.QueryRow(ctx, query, username, passwordHash, fullName, hireDate).Scan(&managerID)
	if err != nil {
		return "", fmt.Errorf("ошибка вызова хранимой функции create_manager: %v", err)
	}

	// Логирование действия
	err = db.LogAction(ctx, managerID, fmt.Sprintf("Менеджер %s был создан", username))
	if err != nil {
		return "", fmt.Errorf("ошибка записи лога: %v", err)
	}

	// Возвращаем ID нового менеджера
	return managerID, nil
}

func (db *db) DeleteAdmin(ctx context.Context, adminID string) error {
	// SQL-запрос для вызова хранимой функции delete_admin
	query := `SELECT delete_admin($1)`

	// Выполнение запроса для вызова хранимой функции
	_, err := db.Pool.Exec(ctx, query, adminID)
	if err != nil {
		return fmt.Errorf("ошибка вызова хранимой функции delete_admin: %v", err)
	}

	// Логирование действия
	err = db.LogAction(ctx, adminID, "Администратор был логически удален")
	if err != nil {
		return fmt.Errorf("ошибка записи лога: %v", err)
	}

	return nil
}

func (db *db) DeleteClient(ctx context.Context, clientID string) error {
	// SQL-запрос для вызова хранимой функции delete_client
	query := `SELECT delete_client($1)`

	// Выполнение запроса для вызова хранимой функции
	_, err := db.Pool.Exec(ctx, query, clientID)
	if err != nil {
		return fmt.Errorf("ошибка вызова хранимой функции delete_client: %v", err)
	}

	// Логирование действия
	err = db.LogAction(ctx, clientID, "Клиент был логически удален")
	if err != nil {
		return fmt.Errorf("ошибка записи лога: %v", err)
	}

	return nil
}

func (db *db) DeleteManager(ctx context.Context, managerID string) error {
	// SQL-запрос для вызова хранимой функции delete_manager
	query := `SELECT delete_manager($1)`

	// Выполнение запроса для вызова хранимой функции
	_, err := db.Pool.Exec(ctx, query, managerID)
	if err != nil {
		return fmt.Errorf("ошибка вызова хранимой функции delete_manager: %v", err)
	}

	// Логирование действия
	err = db.LogAction(ctx, managerID, "Менеджер был логически удален")
	if err != nil {
		return fmt.Errorf("ошибка записи лога: %v", err)
	}

	return nil
}
