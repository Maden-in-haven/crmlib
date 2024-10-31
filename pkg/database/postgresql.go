package database

import (
	"context"
	"fmt"
	"log"
	// "strconv"
	"time"

	"github.com/Maden-in-haven/crmlib/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type db struct {
	Pool *pgxpool.Pool
}

var DB *db

func init() {
	// Загружаем конфигурацию базы данных
	dbConfig := config.LoadDBConfig()

	// Создаем строку подключения
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	// Настраиваем конфигурацию пула соединений
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Ошибка парсинга конфигурации: %v", err)
	}

	// Настраиваем параметры пула
	config.MaxConns = 10                 // Максимум 10 соединений в пуле
	config.MaxConnLifetime = 30 * time.Minute // Максимальное время жизни соединения
	config.HealthCheckPeriod = 1 * time.Minute // Период проверки активности соединений

	// Создаем пул с контекстом и тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Проверяем соединение с помощью Ping
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Ошибка проверки подключения к базе данных (ping): %v", err)
	}

	// Сохраняем пул соединений в глобальной структуре
	DB = &db{Pool: pool}

	log.Println("Успешное подключение к базе данных с использованием пула соединений")
}

func (db *db) LogAction(ctx context.Context, userID, action string) error {
	// SQL-запрос для вставки записи в таблицу логов
	logQuery := `INSERT INTO user_logs (user_id, action) VALUES ($1, $2)`

	// Выполнение SQL-запроса
	_, err := db.Pool.Exec(ctx, logQuery, userID, action)
	if err != nil {
		return fmt.Errorf("ошибка записи лога для пользователя с ID %s: %v", userID, err)
	}

	return nil
}

// 1. CRUD (Create, Read, Update, Delete) операции для каждой сущности:
// Пользователи:

//     CreateUser: Добавление нового пользователя.
//     GetUserByID: Получение пользователя по ID.
//     UpdateUser: Обновление данных пользователя (например, обновление имени пользователя или роли).
//     DeleteUser: Логическое удаление пользователя (помечаем как удаленного, не удаляя физически).

// Администраторы:

//     CreateAdmin: Создание нового администратора с определенными правами.
//     GetAdminByID: Получение администратора по ID.
//     UpdateAdminPermissions: Обновление прав администратора.
//     DeleteAdmin: Логическое удаление администратора.

// Клиенты:

//     CreateClient: Создание нового клиента.
//     GetClientByID: Получение клиента по ID.
//     UpdateClient: Обновление данных клиента (например, номера телефона, имени).
//     DeleteClient: Логическое удаление клиента.

// Менеджеры:

//     CreateManager: Создание нового менеджера.
//     GetManagerByID: Получение менеджера по ID.
//     UpdateManager: Обновление данных менеджера (например, даты приема на работу).
//     DeleteManager: Логическое удаление менеджера.

// 2. Методы для работы с логами пользователей:

//     CreateLog: Создание записи в логах для пользователя (уже вынесено как LogAction).
//     GetUserLogs: Получение логов для определенного пользователя по ID.
//     GetAllLogs: Получение всех логов системы.

// 3. Методы для проверки и управления учетными данными:

//     AuthenticateUser: Аутентификация пользователя (проверка имени пользователя и пароля).
//     ChangePassword: Изменение пароля для пользователя.

// 4. Методы для работы с удаленными пользователями:

//     RestoreUser: Восстановление логически удаленного пользователя.
//     PermanentlyDeleteUser: Полное физическое удаление пользователя (если необходимо).

// 5. Поиск и фильтрация:

//     SearchUsers: Поиск пользователей по имени, роли или другим атрибутам.
//     SearchClients: Поиск клиентов по имени или номеру телефона.
//     SearchManagers: Поиск менеджеров по имени или дате приема на работу.

// 6. Методы для управления правами пользователей:

//     CheckUserPermission: Проверка, имеет ли пользователь определенные права (особенно для администраторов).
//     GrantAdminPermission: Предоставление определенных прав администратору.
//     RevokeAdminPermission: Отзыв прав у администратора.

// 7. Работа с транзакциями (если нужны более сложные операции):

//     BeginTransaction: Начало транзакции.
//     CommitTransaction: Фиксация транзакции.
//     RollbackTransaction: Откат транзакции.

// 8. Работа с метаданными:

//     GetUserCount: Получение общего количества пользователей в системе (включая и исключая удаленных).
//     GetAdminCount: Получение количества администраторов.
//     GetClientCount: Получение количества клиентов.
//     GetManagerCount: Получение количества менеджеров.

// 9. Работа с сессиями (если используется авторизация с сессиями или токенами):

//     CreateSession: Создание сессии для пользователя (например, при аутентификации).
//     GetSession: Получение активной сессии для пользователя.
//     DeleteSession: Завершение сессии (выход пользователя из системы).