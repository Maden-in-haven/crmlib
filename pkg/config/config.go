package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// JWTConfig структура для хранения конфигурации JWT
type JWTConfig struct {
	SecretKey string
}

// GetEnv получает значение переменной окружения или использует значение по умолчанию, если переменная не определена
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Переменная окружения %s не установлена, используется значение по умолчанию: %s", key, defaultValue)
		return defaultValue
	}
	return value
}

// DBConfig структура для хранения конфигурации подключения к базе данных
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// LoadEnv загружает переменные окружения из файла .env
func LoadEnv() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Не удалось загрузить файл .env. Возможно, файл отсутствует или путь к нему некорректный.")
	}
}

// LoadDBConfig загружает конфигурацию для базы данных из переменных окружения
func LoadDBConfig() *DBConfig {
	// Загружаем переменные окружения из файла .env, если они есть
	LoadEnv()

	return &DBConfig{
		Host:     GetEnv("POSTGRESQL_HOST", "localhost"),
		Port:     GetEnv("POSTGRESQL_PORT", "5432"),
		User:     GetEnv("POSTGRESQL_USER", "user"),
		Password: GetEnv("POSTGRESQL_PASSWORD", "password"),
		DBName:   GetEnv("POSTGRESQL_DBNAME", "default_db"),
	}
}

// LoadJWTConfig загружает конфигурацию JWT из переменных окружения
func LoadJWTConfig() *JWTConfig {
	// Загружаем переменные окружения из файла .env, если они есть
	LoadEnv()

	return &JWTConfig{
		SecretKey: GetEnv("JWT_SECRET_KEY", "your_default_secret_key"), // Получаем секретный ключ из переменной окружения
	}
}
