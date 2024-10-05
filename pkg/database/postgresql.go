package database

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Maden-in-haven/crmlib/pkg/config"
	"github.com/jackc/pgx/v5"
)

type db struct {
	Pool *pgx.Conn
}
// глобал
var DB *db

// InitDatabase инициализирует подключение к базе данных
func init() {
	// Загружаем конфигурацию базы данных из переменных окружения
	dbConfig := config.LoadDBConfig()

	// Настраиваем контекст с тайм-аутом для подключения (20 секунд)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Парсинг конфигурации подключения
	config, err := pgx.ParseConfig("")
	if err != nil {
		log.Fatalf("Ошибка парсинга конфигурации: %v", err)
	}

	// Заполняем конфигурацию из переменных окружения
	config.Host = dbConfig.Host
	port, err := strconv.Atoi(dbConfig.Port)
	if err != nil {
		log.Fatalf("Ошибка преобразования порта: %v", err)
	}
	config.Port = uint16(port)
	config.User = dbConfig.User
	config.Password = dbConfig.Password
	config.Database = dbConfig.DBName

	// Подключаемся к базе данных
	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Проверяем подключение с помощью Ping
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("Ошибка проверки подключения к базе данных (ping): %v", err)
	}

	// Сохраняем соединение в глобальной структуре DB
	DB = &db{Pool: conn}

	log.Println("Успешное подключение к базе данных")
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