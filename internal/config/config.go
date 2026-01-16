package config
import "os"


type Config struct {
	Appport string

	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
	}

	JWTSecret string
}


func Load() Config {
	return Config{
		Appport: os.Getenv("APP_PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		DB: struct {
			Host     string
			Port     string
			User     string
			Password string
			Name     string
			SSLMode  string
		}{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),	
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
	}
}