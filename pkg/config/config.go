package config

import (
	"log"
	"os"
	"path/filepath"
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

// func init() {
// 	err := godotenv.Load(findFile(os.Getenv("HOME"), "crm.env"))
// 	if err != nil {
// 		log.Println("Не удалось загрузить файл .env. Возможно, файл отсутствует или путь к нему некорректный.")
// 	}
// }

// LoadDBConfig загружает конфигурацию для базы данных из переменных окружения
func LoadDBConfig() *DBConfig {
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
	return &JWTConfig{
		SecretKey: GetEnv("JWT_SECRET_KEY", "your_default_secret_key"), // Получаем секретный ключ из переменной окружения
	}
}

func findFile(root string, filename string) string {
	var result string

	// Рекурсивно обходим все файлы и папки в директории
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Проверяем, совпадает ли имя файла с искомым
		if info.Name() == filename {
			result = path
			return filepath.SkipDir // Останавливаем поиск после нахождения файла
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Ошибка при поиске файла: %v", err) // Немедленное завершение программы
	}

	if result == "" {
		log.Fatalf("Файл %s не найден", filename) // Немедленное завершение программы
	}

	return result
}