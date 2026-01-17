package mysql

import "doit/urlshortener/pkg/dotenv"

type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewMysqlConfig() *MysqlConfig {
	return &MysqlConfig{
		Host:     dotenv.GetString("DB_HOST", "localhost"),
		Port:     dotenv.GetString("DB_PORT", "3306"),
		User:     dotenv.GetString("DB_USER", "root"),
		Password: dotenv.GetString("DB_PASS", ""),
		Database: dotenv.GetString("DB_NAME", "wec_engagement"),
	}
}
