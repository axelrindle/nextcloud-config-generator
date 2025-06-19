package nextcloud

import (
	"github.com/caarlos0/env/v11"

	_ "github.com/joho/godotenv/autoload"
)

type DatabaseReplica struct {
	User string
	Pass string
	Host string
	Db   string
}

type ConfigDynamic struct {
	AppHost        string   `env:"NC_HOST"`
	AppScheme      string   `env:"NC_SCHEME" envDefault:"https"`
	TrustedDomains []string `env:"NC_TRUSTED_DOMAINS"`

	DatabaseType     string   `env:"NC_DB_TYPE"`
	DatabaseHost     string   `env:"NC_DB_HOST"`
	DatabaseName     string   `env:"NC_DB_NAME"`
	DatabaseUser     string   `env:"NC_DB_USERNAME"`
	DatabasePass     string   `env:"NC_DB_PASSWORD"`
	DatabasePrefix   string   `env:"NC_DB_PREFIX" envDefault:"oc_"`
	DatabaseReplicas []string `env:"NC_DB_REPLICAS" envSeparator:";"`

	MailDomain      string `env:"NC_MAIL_DOMAIN"`
	MailFromAddress string `env:"NC_MAIL_FROM_ADDRESS"`
	MailMode        string `env:"NC_MAIL_MODE" envDefault:"smtp"`
	MailHost        string `env:"NC_MAIL_HOST"`
	MailPort        int16  `env:"NC_MAIL_PORT" envDefault:"25"`
	MailSecure      bool   `env:"NC_MAIL_SECURE" envDefault:"true"`
	MailUser        string `env:"NC_MAIL_USERNAME"`
	MailPass        string `env:"NC_MAIL_PASSWORD"`

	RedisEnabled  bool   `env:"NC_REDIS_ENABLED" envDefault:"false"`
	RedisHost     string `env:"NC_REDIS_HOST"`
	RedisPort     int16  `env:"NC_REDIS_PORT" envDefault:"6379"`
	RedisTimeout  int16  `env:"NC_REDIS_TIMEOUT"`
	RedisUser     string `env:"NC_REDIS_USERNAME"`
	RedisPassword string `env:"NC_REDIS_PASSWORD"`
	RedisDatabase int16  `env:"NC_REDIS_DATABASE" envDefault:"0"`
}

func (cfg *ConfigDynamic) LoadFromEnv() error {
	return env.Parse(cfg)
}
