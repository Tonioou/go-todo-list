package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Configs struct {
	Postgres struct {
		Url string
	}
	LogLevel    string
	ServiceName string
	JaegerURL   string
}

var configRunOnce sync.Once
var configs *Configs
var config = viper.New()

func GetConfig() *Configs {
	configRunOnce.Do(func() {
		config.SetDefault("service.name", "go-todo-list")
		config.SetDefault("log.level", "info")
		config.SetDefault("postgres.host", "localhost")
		config.SetDefault("postgres.port", 5432)
		config.SetDefault("postgres.username", "todo_app")
		config.SetDefault("postgres.password", "todo_app")
		config.SetDefault("postgres.database", "todo_list")
		config.SetDefault("jaeger.url", "http://localhost:14268/api/traces")
		config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		config.AutomaticEnv()

		databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			config.GetString("postgres.username"),
			config.GetString("postgres.password"),
			config.GetString("postgres.host"),
			config.GetString("postgres.port"),
			config.GetString("postgres.database"),
		)

		configs = &Configs{
			Postgres: struct{ Url string }{
				Url: databaseUrl,
			},
			LogLevel:    config.GetString("log.level"),
			ServiceName: config.GetString("service.name"),
			JaegerURL:   config.GetString("jaeger.url"),
		}
	})

	return configs
}
