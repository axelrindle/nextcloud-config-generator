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
	AppHost        string   `env:"NC_HOST" envDesc:"The primary public-facing URL"`
	AppScheme      string   `env:"NC_SCHEME" envDefault:"https"`
	TrustedDomains []string `env:"NC_TRUSTED_DOMAINS" envDesc:"All URLs the nextcloud instance will be accessible at"`

	DatabaseType     string   `env:"NC_DB_TYPE" doc:"dbtype"`
	DatabaseHost     string   `env:"NC_DB_HOST" doc:"dbhost"`
	DatabaseName     string   `env:"NC_DB_NAME" doc:"dbname"`
	DatabaseUser     string   `env:"NC_DB_USERNAME" doc:"dbuser"`
	DatabasePass     string   `env:"NC_DB_PASSWORD" doc:"dbpassword"`
	DatabasePrefix   string   `env:"NC_DB_PREFIX" envDefault:"nc_" doc:"dbtableprefix"`
	DatabaseReplicas []string `env:"NC_DB_REPLICAS" envSeparator:";" doc:"dbreplica"`

	MailDomain      string `env:"NC_MAIL_DOMAIN" doc:"mail_domain"`
	MailFromAddress string `env:"NC_MAIL_FROM_ADDRESS" doc:"mail_from_address"`
	MailMode        string `env:"NC_MAIL_MODE" envDefault:"smtp" doc:"mail_smtpmode"`
	MailHost        string `env:"NC_MAIL_HOST" doc:"mail_smtphost"`
	MailPort        int16  `env:"NC_MAIL_PORT" envDefault:"25" doc:"mail_smtpport"`
	MailSecure      bool   `env:"NC_MAIL_SECURE" envDefault:"true" doc:"mail_smtpsecure"`
	MailUser        string `env:"NC_MAIL_USERNAME" doc:"mail_smtpname"`
	MailPass        string `env:"NC_MAIL_PASSWORD" doc:"mail_smtppassword"`

	RedisEnabled  bool   `env:"NC_REDIS_ENABLED" envDefault:"false" doc:"redis"`
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
