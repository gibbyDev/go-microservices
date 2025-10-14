package config

import (
	"os"
	"strconv"
)

type Env struct {
	Port           int
	JWTSecret      string
	TokenDuration  int
	DatabaseURL    string
	RedisAddr      string
	RedisPassword  string
	RedisDB        int
	EmailHost      string
	EmailPort      int
	EmailUsername  string
	EmailPassword  string
	EmailFrom      string
	FrontendURL    string
}