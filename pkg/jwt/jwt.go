package services

import (
	"github.com/Maden-in-haven/crmlib/pkg/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"errors"
)

// GenerateJWT генерирует JWT токен для указанного пользователя
func GenerateJWT(userID string) (string, error) {
	// Определяем время истечения токена (например, 12 часов)
	tokenExpirationTime := time.Now().Add(12 * time.Hour)

	// Создаем claims с добавлением читаемого времени и типа токена
	claims := jwt.MapClaims{
		"sub":          userID,                                   // ID пользователя
		"exp":          tokenExpirationTime.Unix(),               // Время истечения в Unix формате (обязательно для JWT)
		"iat":          time.Now().Unix(),                        // Время создания в Unix формате
		"typ":          "access",                                 // Тип токена: access (основной токен)
		"exp_readable": tokenExpirationTime.Format(time.RFC3339), // Читаемое время истечения (ISO 8601)
		"iat_readable": time.Now().Format(time.RFC3339),          // Читаемое время создания (ISO 8601)
	}

	// Создаем новый токен с алгоритмом подписи и claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с помощью секретного ключа
	tokenString, err := token.SignedString([]byte(config.LoadJWTConfig().SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


// GenerateRefreshToken генерирует рефреш токен для указанного пользователя
func GenerateRefreshToken(userID string) (string, error) {
	// Определяем время истечения рефреш токена (например, 7 дней)
	tokenExpirationTime := time.Now().Add(7 * 24 * time.Hour)

	// Создаем claims для рефреш токена
	claims := jwt.MapClaims{
		"sub":          userID,                                   // ID пользователя
		"exp":          tokenExpirationTime.Unix(),               // Время истечения в Unix формате
		"iat":          time.Now().Unix(),                        // Время создания в Unix формате
		"typ":          "refresh",                                // Указываем тип токена как "refresh"
		"exp_readable": tokenExpirationTime.Format(time.RFC3339), // Читаемое время истечения
		"iat_readable": time.Now().Format(time.RFC3339),          // Читаемое время создания
	}

	// Создаем новый рефреш токен с алгоритмом подписи и claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем рефреш токен с помощью секретного ключа
	tokenString, err := token.SignedString([]byte(config.LoadJWTConfig().SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Валидация JWT с использованием конфигурации
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// Парсим и валидируем токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неожиданный метод подписи")
		}
		return []byte(config.LoadJWTConfig().SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Возвращаем claims, если токен валиден
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("недействительный токен")
}
