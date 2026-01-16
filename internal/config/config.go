package config
import "os"


type Config struct {
	AppPort string

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
	var cfg Config
	cfg.AppPort = os.Getenv("APP_PORT")
	cfg.JWTSecret = os.Getenv("JWT_SECRET")
	cfg.DB.Host = os.Getenv("PG_HOST")
	cfg.DB.Port = os.Getenv("PG_PORT")
	cfg.DB.User = os.Getenv("PG_USER")
	cfg.DB.Password = os.Getenv("PG_PASSWORD")
	cfg.DB.Name = os.Getenv("PG_DB")
	cfg.DB.SSLMode = os.Getenv("PG_SSL")
	return cfg
}