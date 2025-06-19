package cli

import (
	"log"

	"github.com/axelrindle/nc-cfg-gen/nextcloud"
	"github.com/axelrindle/nc-cfg-gen/print"
	"github.com/spf13/cobra"
)

var dynamicCmd = &cobra.Command{
	Use:     "dynamic",
	Short:   "Generates the dynamic configuration, which may change during starts.",
	GroupID: "generate",
	Run: func(cmd *cobra.Command, args []string) {
		config := &nextcloud.ConfigDynamic{}

		err := config.LoadFromEnv()
		if err != nil {
			log.Fatal(err)
		}

		printer := print.Printer{}

		printer.PrintHead()
		printer.PrintString(config.AppHost, "overwrite.cli.url")
		printer.PrintString(config.AppHost, "overwritehost")
		printer.PrintString(config.AppScheme, "overwriteprotocol")
		printer.PrintStringSlice(config.TrustedDomains, "trusted_domains")

		printer.PrintBool(true, "upgrade.disable-web")
		printer.PrintString("/", "htaccess.RewriteBase")

		printer.PrintString(config.DatabaseType, "dbtype")
		printer.PrintString(config.DatabaseHost, "dbhost")
		printer.PrintString(config.DatabaseName, "dbname")
		printer.PrintString(config.DatabaseUser, "dbuser")
		printer.PrintString(config.DatabasePass, "dbpassword")
		printer.PrintString(config.DatabasePrefix, "dbtableprefix")
		printer.PrintStringMap(config.DatabaseReplicas, "dbreplica")

		printer.PrintString(config.MailDomain, "mail_domain")
		printer.PrintString(config.MailFromAddress, "mail_from_address")
		printer.PrintString(config.MailMode, "mail_smtpmode")
		printer.PrintString(config.MailHost, "mail_smtphost")
		printer.PrintInt16(config.MailPort, "mail_smtpport")
		printer.PrintBool(config.MailSecure, "mail_smtpsecure")
		printer.PrintString(config.MailUser, "mail_smtpname")
		printer.PrintString(config.MailPass, "mail_smtppassword")

		if config.RedisEnabled {
			printer.PrintString("\\OC\\Memcache\\Redis", "memcache.local")
			printer.PrintString("\\OC\\Memcache\\Redis", "memcache.distributed")
			printer.PrintString("\\OC\\Memcache\\Redis", "memcache.locking")

			printer.StartArray("redis")
			printer.PrintString(config.RedisHost, "host")
			printer.PrintInt16(config.RedisPort, "port")
			printer.PrintInt16(config.RedisTimeout, "timeout")
			printer.PrintInt16(config.RedisTimeout, "read_timeout")
			printer.PrintString(config.RedisUser, "user")
			printer.PrintString(config.RedisPassword, "password")
			printer.PrintInt16(config.RedisDatabase, "dbindex")
			printer.EndArray()
		}

		printer.PrintFoot()
	},
}

func init() {
	rootCmd.AddCommand(dynamicCmd)
}
