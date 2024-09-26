package database

import (
	"github.com/Maden-in-haven/crmlib/pkg/config"
	"context"
	"log"
	"strconv"
	"time"
	"github.com/jackc/pgx/v5"
)

// Глобальная переменная для хранения пула соединений
var DbPool *pgx.Conn

// InitDatabase инициализирует подключение к базе данных и сохраняет его в глобальной переменной
func InitDatabase() error {
	// Загружаем конфигурацию базы данных из переменных окружения
	dbConfig := config.LoadDBConfig()

	// Настраиваем контекст с тайм-аутом для подключения (5 секунд)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	config, _ := pgx.ParseConfig("")
	
	config.Host = dbConfig.Host
	port, _ := strconv.Atoi(dbConfig.Port)
	config.Port = uint16(port)
	config.User = dbConfig.User
	config.Password = dbConfig.Password
	config.Database = dbConfig.DBName

	// Подключаемся к базе данных
	dbpool, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		log.Printf("Ошибка подключения к базе данных: %v", err)
		return err
	}

	// Проверяем, что подключение успешно
	err = dbpool.Ping(ctx)
	if err != nil {
		log.Printf("Ошибка проверки подключения к базе данных (ping): %v", err)
		return err
	}

	// Сохраняем пул соединений в глобальную переменную
	DbPool = dbpool

	log.Println("Успешно подключено к базе данных")
	return nil
}
