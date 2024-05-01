package redisqueue

type Config struct {
	UserName	string	`json:"user_name"`
	Host     string `env:"REDIS_HOST,notEmpty"`
	Port     string `env:"REDIS_PORT,notEmpty"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int  `env:"REDIS_DB,notEmpty"`
}
