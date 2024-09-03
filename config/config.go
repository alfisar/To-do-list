package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"todolist/database"
	"todolist/internal/jwthandler"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type directory struct {
	DirUser string
}

type expTime struct {
	ExpLogin time.Duration
}

type hashing struct {
	Key string
}
type Config struct {
	DBSql   *gorm.DB
	Jwt     jwthandler.JwtHandler
	DBRedis map[string]*redis.Client
	DIR     directory
	ExpData expTime
	Hashing hashing
}

var (
	AppConfig *Config
	configErr error
)

func init() {

	var (
		DBSql      *gorm.DB
		jwtSecret  string
		dbRedis    map[string]*redis.Client
		expLogin   time.Duration
		hashingKey string
		err        error
	)

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	fmt.Println("DB_USE : " + os.Getenv("DB_USE"))
	switch os.Getenv("DB_USE") {
	case "sql":
		DBSql, err = database.NewConnSQL()
		if err != nil {
			configErr = err
			return
		}
	default:
		configErr = fmt.Errorf("error DB Use")
		return
	}

	dbRedis, err = database.NewDatabaseRedis()
	if err != nil {
		configErr = err
		return
	}

	dirUser := os.Getenv("DIR_USER")

	if os.Getenv("JWT_SECRET") != "" {
		jwtSecret = os.Getenv("JWT_SECRET")
	}

	if os.Getenv("EXP_LOGIN") != "" {
		dataExp, _ := strconv.Atoi(os.Getenv("EXP_LOGIN"))

		expLogin = time.Minute * time.Duration(dataExp)
	}

	if os.Getenv("HASHING_KEY") != "" {
		hashingKey = os.Getenv("HASHING_KEY")
	}

	AppConfig = &Config{
		DBSql:   DBSql,
		DBRedis: dbRedis,
		Jwt: jwthandler.JwtHandler{
			Secret: jwtSecret,
		},
		DIR: directory{
			DirUser: dirUser,
		},
		ExpData: expTime{
			ExpLogin: expLogin,
		},
		Hashing: hashing{
			Key: hashingKey,
		},
	}

}

func InitConfig(ctx context.Context) (*Config, error) {
	return AppConfig, configErr
}
